package dao

import (
	"database/sql"

	. "github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/table"
	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	. "github.com/go-jet/jet/v2/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func AddNewForgottenPasswordForUser(attempt *model.ForgottenPasswordAttempt) error {
    stmt := ForgottenPasswordAttempt.INSERT(
        ForgottenPasswordAttempt.DateDemande,
        ForgottenPasswordAttempt.LienChangement,
        ForgottenPasswordAttempt.UtilisateurID,
        ).MODEL(attempt)
	_, err := stmt.Exec(db)
    return err
}

func FindPasswordAttemptByUUID(uuidBytes []byte) (*model.Utilisateur, error) {
    user := model.Utilisateur{}
    stmt := SELECT(ForgottenPasswordAttempt.UtilisateurID.AS("utilisateur.id")).FROM(
        ForgottenPasswordAttempt,
        ).WHERE(ForgottenPasswordAttempt.Active.AND(
            ForgottenPasswordAttempt.LienChangement.EQ(String(string(uuidBytes))),
            ))
    return &user, stmt.Query(db, &user)
}

func ChangePassword(password string, uuid string, tx *sql.Tx) error {
    stmt := RawStatement(`
        UPDATE utilisateur SET utilisateur.Password = #password
        WHERE utilisateur.ID = (SELECT UtilisateurId FROM forgotten_password_attempt WHERE LienChangement = #uuid)
        `,
        RawArgs{"#password": password, "#uuid": uuid})
    _, err := stmt.Exec(tx)
    return err
}

func SetPasswordAttemptsInvalid(uuid string, tx *sql.Tx) error {
    stmt := RawStatement(`
        UPDATE forgotten_password_attempt SET forgotten_password_attempt.Active = 0
        WHERE UtilisateurID = (
            SELECT DISTINCT(u.ID) FROM utilisateur u
            LEFT JOIN forgotten_password_attempt fpa ON u.ID = fpa.UtilisateurID
            WHERE fpa.LienChangement = #uuid
        )
        `,
        RawArgs{"#uuid": uuid})
    _, err := stmt.Exec(tx)
    return err
}
