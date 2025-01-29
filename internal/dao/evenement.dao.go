package dao

import (
	"fmt"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	. "github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/table"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	. "github.com/go-jet/jet/v2/mysql"
)

func GetEvenementsAVenir() ([]*model.Evenement, error) {
	// Build the SQL query
	var events []*model.Evenement
	query := SELECT(Evenement.AllColumns).
		FROM(Evenement).
		WHERE(Evenement.Date.GT_EQ(CURRENT_TIMESTAMP())).
        ORDER_BY(Evenement.Date.ASC())

    fmt.Printf("query.DebugSql(): %v\n", query.DebugSql())
	if err := query.Query(db, &events); err != nil {
        fmt.Printf("err: %v\n", err)
		return nil, err
	}

	return events, nil
}

func GetEvenements() ([]*model.Evenement, error) {
    var events []*model.Evenement
    return events, SELECT(Evenement.AllColumns).
		FROM(Evenement).
		WHERE(Evenement.Date.GT_EQ(CURRENT_TIMESTAMP())).
        ORDER_BY(Evenement.Date.ASC()).
        Query(db, &events)
}

func InsertEvenement(evt *model.Evenement) (*model.Evenement, error) {
    _, err := Evenement.INSERT(Evenement.AllColumns).MODEL(evt).Exec(db)
    return evt, err
}

func DeleteEvenement(id int64) error {
    _, err := Evenement.DELETE().WHERE(Evenement.ID.EQ(Int(id))).Exec(db)
    return err
}

func GetEvtById(id int64) (*model.Evenement, error) {
    evt := model.Evenement{}
    stmt := SELECT(Evenement.ID, Evenement.Description, Evenement.Title, Evenement.Date).
        FROM(Evenement).
        WHERE(Evenement.ID.EQ(Int(id)))
    return &evt, stmt.Query(db, &evt)
}

func ModifyEvent(dto *dto.EvenementDto, evt_id int64) error {

    stmt := Evenement.UPDATE().
        SET(
            Evenement.Title.SET(String(dto.Title)),
            Evenement.Description.SET(String(*dto.Description)),
            Evenement.Date.SET(DateTimeT(dto.Date)),
        ).
        WHERE(Evenement.ID.EQ(Int(evt_id)))
    _, err := stmt.Exec(db)
    return err
}
