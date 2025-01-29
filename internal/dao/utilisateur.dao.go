package dao

import (
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	. "github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/table"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	. "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

type Count struct {
    Count int64
}

func GetUsers() ([]model.Utilisateur, error) {
	var users []model.Utilisateur
	stmt := SELECT(Utilisateur.ID, Utilisateur.Prenom, Utilisateur.Nom).FROM(Utilisateur)
	return users, stmt.Query(db, &users)
}

func EmailExists(email string, userId int64) (bool, error) {
	var count Count

	stmt := SELECT(
		COUNT(Utilisateur.ID).AS("count.count"),
	).FROM(
		Utilisateur,
	).WHERE(
		Utilisateur.Email.EQ(String(email)).AND(Utilisateur.ID.NOT_EQ(Int(userId))),
	)
	err := stmt.Query(db, &count)

	if err != nil {
		return false, err
	}
	return count.Count > 0, err
}

func PersistUtilisateurSignup(utilisateur *model.Utilisateur) (*model.Utilisateur, error) {
	stmt := Utilisateur.INSERT(
		Utilisateur.Prenom,
		Utilisateur.Nom,
		Utilisateur.DateCreation,
		Utilisateur.Password,
		Utilisateur.Email,
		Utilisateur.ActivationLink,
		Utilisateur.Role,
	).MODEL(utilisateur)
	return returning(stmt, db)
}

func returning(stmt Statement, tx qrm.DB) (*model.Utilisateur, error) {
	r, err := stmt.Exec(tx)
	if err != nil {
		return nil, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	return FindUserByIdTX(id, tx)
}

func FindUserById(userId int64) (*model.Utilisateur, error) {
	return FindUserByIdTX(userId, db)
}

func FindUserByIdTX(id int64, tx qrm.DB) (*model.Utilisateur, error) {
	var utilisateur model.Utilisateur
	stmt := SELECT(
		Utilisateur.AllColumns.Except(Utilisateur.Password)).FROM(
		Utilisateur).WHERE(
		Utilisateur.ID.EQ(Int(id)),
	)
	return &utilisateur, stmt.Query(tx, &utilisateur)
}

func FindByEmail(email string) (*model.Utilisateur, error) {
	var utilisateur model.Utilisateur
	stmt := SELECT(
		Utilisateur.AllColumns.Except(Utilisateur.ActivationLink)).FROM(
		Utilisateur).WHERE(
		Utilisateur.Email.EQ(String(email)),
	)
	return &utilisateur, stmt.Query(db, &utilisateur)
}

func FindByActivationLink(uuid []byte) (*model.Utilisateur, error) {
	var utilisateur model.Utilisateur
	stmt := SELECT(
		Utilisateur.AllColumns).FROM(
		Utilisateur).WHERE(
		Utilisateur.ActivationLink.EQ(String(string(uuid))),
	)
	return &utilisateur, stmt.Query(db, &utilisateur)
}

func SetActivated(user *model.Utilisateur) (*model.Utilisateur, error) {

	stmt := RawStatement(`
        UPDATE utilisateur SET ActivationLink = NULL, EmailVerified = 1
        WHERE utilisateur.ID = #id
        `,
		RawArgs{"#id": int64(user.ID)},
	)

	_, err := stmt.Exec(db)
	if err != nil {
		return nil, err
	}
	return FindUserByIdTX(int64(user.ID), db)
}

func FindBasicUserByEmail(email string) (*model.Utilisateur, error) {
	var utilisateur model.Utilisateur
	stmt := SELECT(
		Utilisateur.ID,
		Utilisateur.Nom,
		Utilisateur.Prenom,
		Utilisateur.Email,
	).FROM(
		Utilisateur).WHERE(
		Utilisateur.Email.EQ(String(email)),
	)
	return &utilisateur, stmt.Query(db, &utilisateur)
}

func CheckAttemptInLastFiveMinutes(email string, t time.Time) error {
	var utilisateur model.Utilisateur
	stmt := SELECT(ForgottenPasswordAttempt.ID).FROM(
		ForgottenPasswordAttempt.INNER_JOIN(Utilisateur, Utilisateur.ID.EQ(ForgottenPasswordAttempt.UtilisateurID)),
	).WHERE(Utilisateur.Email.EQ(String(email)).AND(
		ForgottenPasswordAttempt.DateDemande.GT(Timestamp(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())),
	))
	return stmt.Query(db, &utilisateur)
}

func MailSameAsOld(email string, userId int64) (bool, error) {
	var count Count
	stmt := SELECT(COUNT(Utilisateur.ID).AS("count.count")).
		FROM(Utilisateur).
		WHERE(Utilisateur.ID.EQ(Int(userId)).
			AND(Utilisateur.Email.EQ(String(email))))
	err := stmt.Query(db, &count)
	if err != nil {
		return false, err
	}
	return count.Count > 0, nil
}

func PersistUtilisateurChange(user *dto.UserChangeDto, userId int64, binary []byte, sameMail bool) (*model.Utilisateur, error) {
	stmt := Utilisateur.UPDATE().SET(
		Utilisateur.Prenom.SET(String(*user.Prenom)),
		Utilisateur.Nom.SET(String(*user.Nom)),
	)
	stmt.WHERE(Utilisateur.ID.EQ(Int(userId)))

	_, err := stmt.Exec(db)
	if err != nil {
		return nil, err
	}
	return FindUserById(userId)
}

func GetUtilisateurs() ([]*model.Utilisateur, error) {
    var users []*model.Utilisateur
    return users, SELECT(Utilisateur.AllColumns.Except(Utilisateur.Password, Utilisateur.ActivationLink)).
        FROM(Utilisateur).
        Query(db, &users)
}
