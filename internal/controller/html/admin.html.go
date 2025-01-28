package html

import (
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/Cyber-cicco/jardin-pc/internal/middleware"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
	"github.com/Cyber-cicco/jardin-pc/internal/views/admin"
	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(r_no_auth, r_auth *gin.RouterGroup) {
	r_no_auth.GET("/admin", LoginPage)
	r_no_auth.POST("/admin", Login)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "", admin.Login())
}

func setAuthCookie(c *gin.Context, jwt string) {
	c.SetCookie(
		"AUTH_TOKEN",
		jwt,
		86000,
		"/admin",
		"",
		true,
		true,
	)
}

func Login(c *gin.Context) {
	var login dto.LoginDto
	err := c.Bind(&login)
	err_map := make(map[string]string)
	if err != nil {
		err_map["main"] = "Format de la requête invalide"
		c.HTML(http.StatusOK, "", admin.LoginForm(err_map))
		return
	}

	machine_infos := c.MustGet(middleware.MachineKey).(dto.RequestMachineInfos)
	auth, diags := service.Login(&login, &machine_infos)

	if diags.IsNotEmpty() {
		err_map = diags.Errors
		c.HTML(http.StatusOK, "", admin.LoginForm(err_map))
		return
	}

	jwt, diags := service.BuildJWTToken(auth)

	if diags != nil {
		err_map = diags.Errors
		c.HTML(http.StatusOK, "", admin.LoginForm(err_map))
		return
	}

	setAuthCookie(c, jwt)

    c.Header("HX-Location", "/admin/events")

    c.HTML(http.StatusOK, "", admin.Login())

}
