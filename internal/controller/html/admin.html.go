package html

import (
	"fmt"
	"net/http"

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
	"github.com/Cyber-cicco/jardin-pc/internal/middleware"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
	"github.com/Cyber-cicco/jardin-pc/internal/views/admin"
	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(r_no_auth, r_auth *gin.RouterGroup) {
	r_no_auth.GET("/admin", LoginPage)
	r_no_auth.POST("/admin", Login)
	r_auth.GET("/admin/events", EvenementsDashboard)
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
		c.HTML(http.StatusBadRequest, "", admin.LoginForm(err_map))
		return
	}

	jwt, diags := service.BuildJWTToken(auth)

	if diags.IsNotEmpty() {
		err_map = diags.Errors
        fmt.Printf("err_map: %v\n", err_map)
		c.HTML(http.StatusBadRequest, "", admin.LoginForm(err_map))
		return
	}

	setAuthCookie(c, jwt)
    c.Header("HX-Retarget", "body")
    c.Header("HX-Redirect", "/admin/events")
    c.HTML(http.StatusFound, "", admin.Login())

}

func EvenementsDashboard(c *gin.Context) {

    before, after, err := service.GetEvenements()

    if err != nil {
        before = []*model.Evenement{}
        after = []*model.Evenement{}
    }

    c.HTML(http.StatusOK, "", admin.EvenementDashBoard(before, after))
}

func AddEvenement(c *gin.Context) {

    var evt model.Evenement
    auth := c.MustGet(config.AuthKey).(dto.AuthDto)

    err_map := make(map[string]string)
    value_map := make(map[string]string)
    err := c.Bind(&evt)
    if err != nil {
        c.Header("HX-Retarget", "form")
        c.HTML(http.StatusOK, "", admin.AddEvtForm(err_map, value_map))
    }

    value_map["title"] = evt.Title
    if evt.Description != nil {
        value_map["description"] = *evt.Description
    }
    value_map["date"] = evt.Date.String()

    service.AddEvenement(int64(auth.Id), &evt)
}
