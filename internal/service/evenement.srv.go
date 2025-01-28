package service

import (
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
)

func GetEvenementsAVenir() ([]*model.Evenement, error) {
    return dao.GetEvenementsAVenir()
}

func GetEvenements() ([]*model.Evenement, []*model.Evenement, error) {
    evts, err := dao.GetEvenements()
    if err != nil {
        return nil, nil, err
    }
    now := time.Now()
    before := []*model.Evenement{}
    after := []*model.Evenement{}
    for _, evt := range evts {
        if evt.Date.After(now) {
            after = append(after, evt)
        } else {
            before = append(before, evt)
        }
    }
    return before, after, nil
}

