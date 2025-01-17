package main

import (
	"fmt"
	"log"
	"reservation-service/internal/server"
)

func main() {
	cfg := server.NewConfig()
	srv := server.NewServer(cfg)

	fmt.Println("server started at the port:", cfg.Port)
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}

}
