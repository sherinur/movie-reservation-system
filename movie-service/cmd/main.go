package main

import (
	"log"

	"movie-service/internal/server"
)

func main() {
	cfg := server.GetConfig()
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		log.Fatal(err)
	}
}
