package service

import (
	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
)

func GetEvenements() ([]*model.Evenement, error) {
    return dao.GetEvenementsAVenir()
}
