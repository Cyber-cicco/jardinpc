package dto

import (
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/google/uuid"
	"github.com/mileusna/useragent"
)

type Roles []string

var (
	DEFAULT_ROLES = Roles([]string{"COLLAB"})
)

type SignupDto struct {
	Identifiant          *string `json:"identifier"`
	Prenom               *string `json:"prenom"`
	Nom                  *string `json:"nom"`
	Email                *string `json:"email"`
	Password             *string `json:"password"`
	NumTel               *string `json:"numTel"`
	PasswordConfirmation *string `json:"passwordConf"`
}

type SignuResponseDto struct {
	*model.Utilisateur
	Jwt string `json:"jwt"`
}

type UserChangeDto struct {
	Identifiant *string `json:"Identifiant"`
	Prenom      *string `json:"Prenom"`
	Nom         *string `json:"Nom"`
	Email       *string `json:"Email"`
	NumTel      *string `json:"NumTel"`
}

type AuthDto struct {
	Exp      float64
	Id       float64
	Roles    []string
	UserName string
	Verified bool
}

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type PasswordRecoveryDto struct {
	Email  string
	Nom    string
	Prenom string
	Uuid   uuid.UUID
}

type PasswordChangeDto struct {
	Password     string `form:"password"`
	PasswordConf string `form:"password-conf"`
}

type GroupsDto struct {
	Members     []*SimpleColocDto
	Invitations []*InviteDto
}

type SimpleColocDto struct {
	ID    int64
	Label string
}

type InviteDto struct {
	ID     int64
	Label  string
	Nom    string
	Prenom string
	NumTel string
}

type RequestMachineInfos struct {
	useragent.UserAgent
	Date     time.Time
	IpAdress string
}
