package validator

import (
	"net/http"
	"net/mail"
	"regexp"
	"unicode"

	"github.com/Cyber-cicco/jardin-pc/internal/dao"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
)

func ValidateUserChange(user *dto.UserChangeDto, userId int64) *Diagnostics {
	diags := GetDiagnostics(http.StatusBadRequest)

	// Email validation
	diags.PushIfNullOrBlank(user.Email, "email", "Le mail ne doit pas être vide")

	// Check if account already exists
	if user.Email != nil {
		exists, err := dao.EmailExists(*user.Email, userId)
		if err != nil {
			diags = GetDiagnostics(http.StatusInternalServerError)
			diags.AppendError("main", "Erreur dans la vérification de l'existence du mail")
			return diags
		}
		diags.PushIfConditionIsTrue(exists, "email", "Le mail appartient déjà à un autre utilisateur")

	}
	diags.PushIfLenAbove(255, user.Email, "email", "Le mail est trop long")

	_, err := mail.ParseAddress(*user.Email)
	diags.PushIfConditionIsTrue(err != nil, "email", "Le mail n'est pas valide")

	// Prenom validation
	diags.PushIfNullOrBlank(user.Nom, "nom", "Le nom ne peut pas être vide")
	diags.PushIfLenAbove(120, user.Nom, "nom", "Le nom est trop long")

	// Nom validation
	diags.PushIfNullOrBlank(user.Prenom, "prenom", "Le prénom ne peut pas être vide")
	diags.PushIfLenAbove(120, user.Prenom, "prenom", "Le prénom est trop long")

	return diags
}

func ValidateSignup(signup *dto.SignupDto) *Diagnostics {
    diags := ValidateUserChange(&dto.UserChangeDto{
    	Identifiant: signup.Identifiant,
    	Prenom:      signup.Prenom,
    	Nom:         signup.Nom,
    	Email:       signup.Email,
    	NumTel:      signup.NumTel,
    }, 0)

	ValidatePassword(*signup.Password, *signup.PasswordConfirmation, diags)

	return diags
}

func ValidatePassword(pwd, pwdConf string, d *Diagnostics) {

	if len(pwd) < 8 {
		d.AppendError("password", "Le mot de passe est trop court")
		return
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false
	for _, r := range pwd {
		if unicode.IsDigit(r) {
			hasDigit = true
			continue
		}
		if unicode.IsUpper(r) {
			hasUpperCase = true
			continue
		}
		if unicode.IsLower(r) {
			hasLowerCase = true
			continue
		}
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
			hasSpecialChar = true
		}
	}
	d.PushIfConditionIsTrue(!hasUpperCase, "password", "Le mot de passe doit avoir au moins une lettre en majuscule")
	d.PushIfConditionIsTrue(!hasLowerCase, "password", "Le mot de passe doit avoir au moins une lettre en minuscule")
	d.PushIfConditionIsTrue(!hasDigit, "password", "Le mot de passe doit avoir un nombre")
	d.PushIfConditionIsTrue(!hasSpecialChar, "password", "Le mot de passe doit posséder au moins un caractère spécial")
	//Matching password validation
	d.PushIfConditionIsTrue(pwd != pwdConf, "password", "Les mots de passe ne correspondent pas")
}

func ValidateNumTel(diags *Diagnostics, numTel *string) {

	diags.PushIfNullOrBlank(numTel, "numTel", "Le numéro de téléphone ne peut pas être nul")
	diags.PushIfLenAbove(20, numTel, "numTel", "Le numéro de téléphone est trop long")
	if numTel != nil {
		match := regexp.MustCompile(`^\+((?:9[679]|8[035789]|6[789]|5[90]|42|3[578]|2[1-689])|9[0-58]|8[1246]|6[0-6]|5[1-8]|4[013-9]|3[0-469]|2[70]|7|1)(?:\W*\d){0,13}\d$`)
		diags.PushIfConditionIsTrue(!match.Match([]byte(*numTel)), "numTel", "Le numéro de téléphone est invalide")
	}
}

