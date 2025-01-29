package service

import (
	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
)

func GetUtilisateurs() ([]*model.Utilisateur, error) {
    return dao.GetUtilisateurs()
}
