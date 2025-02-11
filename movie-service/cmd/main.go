package main

import (
	"os"

	"movie-service/internal/server"
)

func main() {
	cfg := server.GetConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
