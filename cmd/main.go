package main

import (
	"fmt"
	"os"
	"udp-map/internal/server"
)

func main() {
	server.StartServer("8080")
	if err := server.StartServer("8080"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
