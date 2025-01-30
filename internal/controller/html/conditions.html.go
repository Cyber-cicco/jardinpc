package html

import (
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/internal/views/conditions"
	"github.com/gin-gonic/gin"
)

func InitConditionsRoutes(r *gin.RouterGroup) {
    r.GET("/conditions", GetConditions)
}

func GetConditions(c *gin.Context) {
    if IsHtmxReq(c) {
        c.HTML(http.StatusOK, "", conditions.ConditionsUtilisationSection())
        return
    }
    c.HTML(http.StatusOK, "", conditions.ConditionsUtilisation())
}
