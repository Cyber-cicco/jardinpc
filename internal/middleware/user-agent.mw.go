package middleware

import (
	"strings"
	"time"

	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
)

const (
    // Clés utilisées dans le contexte pour vérfier l'identité de l'agent
    MachineKey = "machine"
)

// middleware qui extrait les informations de l'IP de l'utilisateur et de sa machine
// pour les ajouter au contexte Gin
func UserInfoMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        // Récupère le User Agent
        userAgent := c.Request.UserAgent()

        // Récupère l'adresse IP de la requête.
        ipAddress := getIPAddress(c)

        // Permet de récupérer les infos du user agent
        ua := useragent.Parse(userAgent)
        rmi := dto.RequestMachineInfos{
            UserAgent: ua,
            Date: time.Now(),
            IpAdress: ipAddress,
        }

        // Ajoute le user agent dans le contexte gin
        c.Set(MachineKey, rmi)
        c.Next()
    }
}

// helper function pour récupérer l'adresse IP dans le header de la requête
// Gère le cas o l'utilisateur se trouve derrière un proxy ou load balancer
func getIPAddress(c *gin.Context) string {
    // Essaie de récupérer l'IP du header X-Forwarded-For
    ipAddress := c.Request.Header.Get("X-Forwarded-For")
    if ipAddress != "" {
        // X-Forwarded-For can peut contenir plusieurs adresses IP
        // On récupère la dernière
        addresses := strings.Split(ipAddress, ",")
        ipAddress = strings.TrimSpace(addresses[len(addresses)-1])
    }

    // Si le header n'est pas présent, on utilise l'adresse Remote
    if ipAddress == "" {
        ipAddress = c.Request.RemoteAddr
    }

    return strings.Split(ipAddress, ":")[0]
}

