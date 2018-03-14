package main

import (
	"fmt"

	"github.com/Barna001/go-auth/config"
	"github.com/Barna001/go-auth/database"
	"github.com/Barna001/go-auth/errors"
	"github.com/Barna001/go-auth/http"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	fmt.Println("hello there general kenobi")
	config.LoadEnv()
	var authConfig config.AuthConfig
	err := envconfig.Process("authapp", &authConfig)
	errors.CriticalHandling(err)
	db := database.TextDB{Location: authConfig.TextDBLocation}
	webServer := http.Server{Port: authConfig.WebserverPort, Db: db}
	webServer.StartServer()
}
