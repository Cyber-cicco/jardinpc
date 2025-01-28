package main

import (
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/controller"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
	"github.com/Cyber-cicco/jardin-pc/internal/service"
)

func main() {
    config.InitConfig()
    dao.InitDB()
    service.NewAttemptMap()
    controller.InitController()
}
