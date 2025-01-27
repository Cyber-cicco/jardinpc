package config

import (
	"encoding/json"
	"os"
	"sync"
)

// Represents every field in the config file
// located in resources directory
type Config struct {
	Database  `json:"database"`
	JWTConfig `json:"jwt"`
}

// Fields for the jwt creation
type JWTConfig struct {
	JWTSecret string `json:"secret"`
}

// Informations for database connection
type Database struct {
	DBUser     string `json:"db-user"`
	DBPassword string `json:"db-password"`
	DBName     string `json:"db-name"`
	DBAdress   string `json:"db-adress"`
}

const (
	AuthKey = "auth"
)

// Checks if there is any reference to env variables
// and changes the struct accordingly
func (c *Config) setupEnv() {
	c.DBUser = os.ExpandEnv(c.DBUser)
	c.DBPassword = os.ExpandEnv(c.DBPassword)
	c.DBName = os.ExpandEnv(c.DBName)
	c.DBAdress = os.ExpandEnv(c.DBAdress)
	c.JWTSecret = os.ExpandEnv(c.JWTSecret)
}

var (
	Conf *Config
	once sync.Once
)

func InitConfig() {
	once.Do(func() {
		file, err := os.ReadFile("../resources/application.json")
		if err != nil {
			panic(err)
		}
		if err = json.Unmarshal(file, &Conf); err != nil {
			panic(err)
		}
		Conf.setupEnv()
	})
}
