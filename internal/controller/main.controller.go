package controller

import (
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/controller/html"
	"github.com/Cyber-cicco/jardin-pc/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitController() {
    
    router := gin.Default()
    router.HTMLRender = &config.TemplRender{}

	//Add static routes
	router.Static("/", "../resources/static/")
	router.StaticFile("/favicon.ico", "../resources/static/favicon.ico")

	//Create the groupes
	baseGroup := router.Group("/")
	authGroup := router.Group("/")
    baseGroup.Use(middleware.UserInfoMiddleware())
    authGroup.Use(middleware.UserInfoMiddleware())
    authGroup.Use(middleware.Authenticate)
    html.InitHomeRoutes(baseGroup)
    html.InitEvenementsRoutes(baseGroup)
    html.InitConditionsRoutes(baseGroup)
    html.InitAdminRoutes(baseGroup, authGroup)
    router.Run("0.0.0.0:8001")

}
