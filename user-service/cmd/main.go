package main

import (
	"os"

	"user-service/configs"
	"user-service/internal/server"
)

func main() {
	cfg := configs.GetConfig()
	if cfg.GoEnv == "test" {
		return
	}

	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
