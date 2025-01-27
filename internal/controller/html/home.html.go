package html

import (
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/internal/views/home"
	"github.com/gin-gonic/gin"
)

func InitHomeRoutes(r *gin.RouterGroup) {
    r.GET("/", getHomePage)
}

func getHomePage(c *gin.Context) {
    c.HTML(http.StatusOK, "", home.Home()) 
}
