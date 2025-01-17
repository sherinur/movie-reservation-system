package main

import (
	"os"

	"user-service/internal/server"
)

func main() {
	cfg := server.GetConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
