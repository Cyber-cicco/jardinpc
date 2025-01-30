package html

import (
	"fmt"
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
	"github.com/Cyber-cicco/jardin-pc/internal/views/home"
	"github.com/gin-gonic/gin"
)

func InitHomeRoutes(r *gin.RouterGroup) {
    r.GET("/", getHomePage)
    r.GET("/home", getHomePage)
}

func getHomePage(c *gin.Context) {
    fmt.Printf("config.Conf.DBName: %v\n", config.Conf.DBName)
    fmt.Printf("config.Conf.DBUser: %v\n", config.Conf.DBUser)
    fmt.Printf("config.Conf.DBPassword: %v\n", config.Conf.DBPassword)
    evts, err := service.GetEvenementsAVenir()
    if err != nil {
        evts = []*model.Evenement{}
    }
    if IsHtmxReq(c) {
        c.HTML(http.StatusOK, "", home.InnerHome(evts)) 
        return
    }
    c.HTML(http.StatusOK, "", home.Home(evts)) 
}
