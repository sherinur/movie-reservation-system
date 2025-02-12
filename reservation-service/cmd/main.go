package main

import (
	"log"

	"reservation-service/internal/server"
)

func main() {
	cfg := server.NewConfig()
	srv := server.NewServer(cfg)

	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
