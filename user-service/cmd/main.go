package main

import (
	"os"

	"user-service/internal/server"

	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

func main() {
	logging.Init()

	cfg := server.GetConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		os.Exit(1)
	}
}
