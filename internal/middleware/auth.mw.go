package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
	"github.com/gin-gonic/gin"
)

// Permet de stopper l'exécution du contexte en indiquant que l'opération n'est pas autorisée
func abortFunc(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort() 
}

// Permet de vérifier l'authentification via le header Authorization
func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	auth, err := service.CheckJWT(tokenString)
	if err != nil {
		fmt.Printf("err in jwt middleware: %v\n", err)
		abortFunc(c)
	}
	c.Set(config.AuthKey, auth)
	c.Next()
}

