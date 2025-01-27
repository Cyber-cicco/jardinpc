package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/Cyber-cicco/jardin-pc/internal/validator"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(signup *dto.SignupDto) (*model.Utilisateur, *validator.Diagnostics) {

	validationDiags := validator.ValidateSignup(signup)
	serverDiags := validator.GetDiagnostics(http.StatusInternalServerError)

	if validationDiags.IsNotEmpty() {
		return nil, validationDiags
	}

	password, err := bcrypt.GenerateFromPassword([]byte(*signup.Password), 14)
	if err != nil {
		serverDiags.AppendError("main", "Le message n'a pas pu être encrypté")
		return nil, serverDiags
	}
	defaultRoles, err := json.Marshal(dto.DEFAULT_ROLES)
	if err != nil {
		serverDiags.AppendError("main", "Erreur lors de la récupération des rôles de l'utilisateur")
		return nil, serverDiags
	}

	uuid := uuid.New()
	binary, _ := uuid.MarshalBinary()
	passStr := string(password)
	utilisateur := &model.Utilisateur{
		Prenom:         *signup.Prenom,
		Nom:            *signup.Nom,
		DateCreation:   time.Now(),
		Password:       &passStr,
		Email:          *signup.Email,
		Role:           string(defaultRoles),
		ActivationLink: &binary,
	}
	utilisateur, err = dao.PersistUtilisateurSignup(utilisateur)
	if err != nil || utilisateur.ID == 0 {
		serverDiags.AppendError("main", "La création du compte a échoué pour une raison liée au serveur de base de données")
        fmt.Printf("err: %v\n", err)
		return nil, serverDiags
	}

	return utilisateur, validationDiags
}

func Login(login *dto.LoginDto, machine *dto.RequestMachineInfos) (*model.Utilisateur, *validator.Diagnostics) {
	diags := validator.GetDiagnostics(http.StatusForbidden)

	// Verify if email exists
	user, err := dao.FindByEmail(login.Email)
	if err != nil {
		diags.AppendError("main", "L'authentification a échoué")
		return nil, diags
	}

	// Check if ip is blocked
	block := Attempts.CheckIfBlocked(login, machine)

	switch block {
	case TEMP_BLOCK:
		{
			Attempts.AddEntry(login, machine)
			diags.AppendError("main", "Vous avez échoué à vous authentifier un trop grand nombre de fois. Veuillez réessayer dans cinq minutes")
			return nil, diags
		}
	case PERMA_BLOCK:
		{
			diags.AppendError("main",  "Vous êtes définitivement interdit d'utilisation de l'application. Veuillez vous rapprocher d'un administrateur")
			return nil, diags
		}
	}

	//Check password
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(login.Password))
	if err != nil {
		diags.AppendError("main", "L'authentification a échoué.")
		Attempts.AddEntry(login, machine)
		return nil, diags
	}
    user.Password = nil
	return user, diags
}

func buildUUIDFromParam(param string) ([]byte, error) {
	uuid, err := uuid.Parse(param)
	if err != nil {
		return nil, err
	}

	return uuid.MarshalBinary()

}

func DoAccountVerification(param string) bool {

	uuidBytes, err := buildUUIDFromParam(param)

	if err != nil {
		return false
	}

	user, err := dao.FindByActivationLink(uuidBytes)

	if err != nil {
		return false
	}

	user, err = dao.SetActivated(user)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
	}

	return err == nil
}

func CheckIfPasswordChangeUUIDExists(param string) bool {

	uuid, err := uuid.Parse(param)
	if err != nil {
		return false
	}

	uuidBytes, err := uuid.MarshalBinary()

	if err != nil {
		return false
	}

	_, err = dao.FindPasswordAttemptByUUID(uuidBytes)

	return err == nil

}

func SetNewPasswordChangeAttempt(auth *dto.LoginDto) (*dto.PasswordRecoveryDto, bool) {
	user, err := dao.FindBasicUserByEmail(auth.Email)
	if err != nil {
		fmt.Printf("An attempt was made to change password on invalid email : %s\n", auth.Email)
		return nil, false
	}
	uuid := uuid.New()
	binary, err := uuid.MarshalBinary()

	if err != nil {
		fmt.Printf("Internal Server error while creating an UUID\n")
		return nil, false
	}
	attempt := &model.ForgottenPasswordAttempt{
		DateDemande:    time.Now(),
		LienChangement: binary,
		UtilisateurID:  user.ID,
	}
	err = dao.AddNewForgottenPasswordForUser(attempt)
	if err != nil {
		fmt.Printf("Internal Server error while adding a ForgottenPasswordAttempt\n")
		return nil, false
	}
	return &dto.PasswordRecoveryDto{
		Email:  user.Email,
		Nom:    user.Nom,
		Prenom: user.Prenom,
		Uuid:   uuid,
	}, true
}

func CheckIfAttemptMade(email string) bool {
	now := time.Now().Add(-time.Minute * 5)
	err := dao.CheckAttemptInLastFiveMinutes(email, now)
	return err == nil
}

func ChangePassword(pswds *dto.PasswordChangeDto, uuidParam string) *validator.Diagnostics {
	diags := validator.GetDiagnostics(http.StatusOK)
	validator.ValidatePassword(pswds.Password, pswds.PasswordConf, diags)
	if diags.IsNotEmpty() {
		return diags
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(pswds.Password), 14)
	uuidBytes, err := buildUUIDFromParam(uuidParam)

	if err != nil {
		diags.AppendError("password", "Erreur interne")
		return diags
	}

	tx := dao.BeginTransaction()
	err = dao.ChangePassword(string(password), string(uuidBytes), tx)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		diags.AppendError("password", "Erreur de la base de donnée lors du changement de mot de passe.")
		tx.Rollback()
		return diags
	}
	err = dao.SetPasswordAttemptsInvalid(string(uuidBytes), tx)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		diags.AppendError("password", "Erreur de la base de donnée lors du changement de mot de passe.")
		tx.Rollback()
		return diags
	}
	tx.Commit()

	return diags
}

