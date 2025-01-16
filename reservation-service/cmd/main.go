package main

import (
	"fmt"
	"log"
	"reservation-service/reservation-service/internal/server"
)

func main() {
	cfg := server.NewConfig("8080")

	srv := server.NewServer(cfg)
	fmt.Println("server started at the port:", cfg.Port)
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}

}
