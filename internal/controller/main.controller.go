package controller

import (
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/controller/html"
	"github.com/gin-gonic/gin"
)

func InitController() {
    
    router := gin.Default()
    router.HTMLRender = &config.TemplRender{}

	//Add static routes
	router.Static("/css", "../resources/static/css")
	router.Static("/js", "../resources/static/js")
	router.Static("/img", "../resources/static/img")
	router.StaticFile("/favicon.ico", "../resources/static/favicon.ico")

	//Create the groupes
	baseGroup := router.Group("/")
    html.InitHomeRoutes(baseGroup)
    html.InitEvenementsRoutes(baseGroup)
    router.Run("0.0.0.0:8001")

}
