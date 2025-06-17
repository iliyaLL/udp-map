package main

import (
	"flag"
	"log"

	"udp-map/internal/server"
)

func main() {
	flag.Parse()

	if err := server.StartServer(portFlag); err != nil {
		log.Fatal("[ERROR]: ", err)
	}
}
