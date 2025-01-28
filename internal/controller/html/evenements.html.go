package html

import (
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
	"github.com/Cyber-cicco/jardin-pc/internal/views/evenements"
	"github.com/gin-gonic/gin"
)

func InitEvenementsRoutes(r *gin.RouterGroup) {
    r.GET("/events", getEventsPage)
}

func getEventsPage(c *gin.Context) {
    evts, err := service.GetEvenementsAVenir()
    if err != nil {
        evts = []*model.Evenement{}
    }
    if IsHtmxReq(c) {
        c.HTML(http.StatusOK, "", evenements.InnerEvenements(evts)) 
        return
    }
    c.HTML(http.StatusOK, "", evenements.Evenements(evts)) 
}
