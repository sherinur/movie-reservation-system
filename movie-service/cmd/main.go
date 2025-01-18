package main

import (
	"movie-service/internal/server"
	"os"
)

func main() {
	cfg := server.NewConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
