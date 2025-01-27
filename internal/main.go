package main

import (
	"github.com/Cyber-cicco/jardin-pc/internal/config"
	"github.com/Cyber-cicco/jardin-pc/internal/controller"
	"github.com/Cyber-cicco/jardin-pc/internal/dao"
)

func main() {
    config.InitConfig()
    dao.InitDB()
    controller.InitController()
}
