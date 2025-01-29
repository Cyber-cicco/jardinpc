package service

import (
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/Cyber-cicco/jardin-pc/internal/validator"
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

func AddEvenement(user_id int64, dto *dto.EvenementDto) (*model.Evenement, *validator.Diagnostics) {
	diags := validator.ValidateEvenement(dto)
	if diags.IsNotEmpty() {
		return nil, diags
	}
    evt := &model.Evenement{
    	Title:        dto.Title,
    	Description:  dto.Description,
    	Date:         dto.Date,
    	DateCreation: time.Now(),
    	CreateurID:   user_id,
    }
    evt, err := dao.InsertEvenement(evt)
	if err != nil {
		diags.AppendError("main", "La création de l'événement a échoué pour une raison liée au serveur")
		return nil, diags
	}
	return evt, diags
}
