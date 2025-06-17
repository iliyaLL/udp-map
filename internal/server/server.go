package server

import (
	"fmt"
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
			fmt.Println("Received command:", command)
			response = kvstorage.Get(command)
		default:
			response = "(error) ERR unknown command"
		}

		conn.WriteToUDP([]byte(response+"\n"), addr)
	}
}
