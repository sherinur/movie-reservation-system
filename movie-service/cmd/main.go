package main

import "flag"

var Port = flag.String("port", "8080", "port number")

func main() {
	flag.Parse()

}
