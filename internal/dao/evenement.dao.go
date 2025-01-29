package dao

import (
	"fmt"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	. "github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/table"
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
