//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Utilisateur struct {
	ID             int64 `sql:"primary_key"`
	Prenom         string
	Nom            string
	Role           string
	DateCreation   time.Time
	Password       *string
	Email          string
	ActivationLink *[]byte
}
