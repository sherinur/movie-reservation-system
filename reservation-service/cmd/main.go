package main

import (
	"log"
	"reservation-service/reservation-service/internal/server"
)

var Port = "8080"

func main() {
	cfg := server.NewConfig(Port)
	serv := server.NewServer(cfg)
	err := serv.Start()
	if err != nil {
		log.Fatal(err)
	}
}
