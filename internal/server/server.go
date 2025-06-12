package server

import (
	"net"
	"strings"
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

	// Read from UDP listener in endless loop
	for {
		buf := make([]byte, 1024)
		_, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return err
		}

		command := strings.TrimSpace(string(buf))
		var response string
		switch {
		case strings.HasPrefix(strings.ToLower(command), "ping"):
			response = "PONG"
		case strings.HasPrefix(strings.ToLower(command), "set"):
			response = kvstorage.SET(command)
		case strings.HasPrefix(strings.ToLower(command), "get"):
			response = kvstorage.GET(command)
		default:
			response = "(error) ERR unknown command"
		}
	}

	return nil
}
