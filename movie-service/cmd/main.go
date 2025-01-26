package main

import (
	"os"

	"movie-service/internal/server"
)

func main() {
	cfg := server.NewConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
