package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/Cyber-cicco/jardin-pc/internal/validator"
	"github.com/golang-jwt/jwt"
)

func CheckJWT(tokenString string) (*dto.AuthDto, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(config.Conf.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid JWT")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Convert roles to a list of string
		rolesI, ok := claims["roles"].([]interface{})
		roles := make([]string, len(rolesI))
		for i, roleI := range rolesI {
			roles[i], ok = roleI.(string)
			if !ok {
				return nil, errors.New("roles could not be parsed from jwt")
			}
		}

		expiration, ok := claims["exp"].(float64)
		if !ok {
            return nil, errors.New("expiration of jwt could not be parsed")
		}
		id, ok := claims["id"].(float64)
		if !ok {
			return nil, errors.New("id of user could not be parsed from JWT")
		}
		username, ok := claims["username"].(string)
		if !ok {
			return nil, errors.New("username of user could not be parsed from JWT")
		}

		verified, ok := claims["verified"].(bool)
		if !ok {
			return nil, errors.New("verification of mail could not be parsed from JWT")
		}

		auth := dto.AuthDto{
			Exp:      expiration,
			Id:       id,
			Roles:    roles,
			UserName: username,
			Verified: verified,
		}
        return &auth, nil
	} 
    return nil, errors.New("could not access claims of jwt")
}

func BuildJWTToken(user *model.Utilisateur) (string, *validator.Diagnostics) {
	diags := validator.GetDiagnostics(http.StatusInternalServerError)
	//Get roles array from string
	var roles []string
	err := json.Unmarshal([]byte(user.Role), &roles)
	if err != nil {
		print("error Unmarshal")
		fmt.Printf("user.Role: %v\n", user.Role)
		diags.AppendError("main", "Could not Unmarshal roles")
		return "", diags
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"roles":    roles,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	tokenStr, err := token.SignedString([]byte(config.Conf.JWTSecret))
	if err != nil {
        fmt.Printf("err: %v\n", err)
		diags := validator.GetDiagnostics(http.StatusInternalServerError)
		diags.AppendError("main", "L'authentification a échoué.")
		return "", diags
	}
	return tokenStr, diags
}


