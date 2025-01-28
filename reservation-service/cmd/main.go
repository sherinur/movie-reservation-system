package main

import (
	"log"

	"reservation-service/internal/server"

	"github.com/sherinur/movie-reservation-system/pkg/logging"
)

func main() {
	logging.Init()

	cfg := server.NewConfig()
	srv := server.NewServer(cfg)

	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
