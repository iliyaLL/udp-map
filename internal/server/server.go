package server

import (
	"log"
	"net"
	"strings"

	"udp-map/pkg/kvstorage"
)

func StartServer(port string) error {
	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return err
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Log the server start
	log.Printf("Server started on port %s", port)
	// Log the server start
	log.Printf("Listening for UDP connections on %s", udpAddr.String())

	// Read from UDP listener in endless loop
	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}

		command := strings.TrimSpace(string(buf[:n]))
		var response string
		switch {
		case strings.HasPrefix(strings.ToLower(command), "ping"):
			response = "PONG"
		case strings.HasPrefix(strings.ToLower(command), "set"):
			response = kvstorage.Set(command)
		case strings.HasPrefix(strings.ToLower(command), "get"):
			response = kvstorage.Get(command)
		default:
			response = "(error) ERR unknown command"
		}

		conn.WriteToUDP([]byte(response+"\n"), addr)
	}
}
