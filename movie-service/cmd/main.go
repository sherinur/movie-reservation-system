package main

import (
	"flag"
	"fmt"

	"movie-service/internal/server"
	"os"
)

var Port = flag.String("port", "8080", "port number")

func main() {
	flag.Parse()

	cfg := server.NewConfig(*Port)
	apiServer := server.NewServer(cfg)

	err := apiServer.Start()
	fmt.Println("Server start on port: ", Port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
