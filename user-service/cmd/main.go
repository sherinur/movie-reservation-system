// @title User Service API
// @version 1.0
// @description This is a user service API for movie reservation system
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"os"

	"user-service/configs"
	"user-service/internal/server"
)

func main() {
	cfg := configs.GetConfig()
	if cfg.GoEnv == "dev" {
		return
	}

	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
