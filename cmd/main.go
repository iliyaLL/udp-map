package main

import (
	"flag"
	"fmt"
	"os"

	"udp-map/internal/server"
)

func main() {
	flag.Parse()

	if err := server.StartServer(portFlag); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
