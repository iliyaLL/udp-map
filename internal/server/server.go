package server

import (
	"net"
)

func StartServer(port string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	return nil
}
