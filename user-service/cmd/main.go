package main

import (
	"flag"
	"fmt"
	"os"

	"user-service/internal/server"
)

var PORT = flag.String("port", "8080", "port number")

func main() {
	flag.Parse()

	cfg := server.NewConfig(":" + *PORT)
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
