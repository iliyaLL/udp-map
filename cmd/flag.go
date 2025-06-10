package main

import "flag"

var (
	// Flags
	portFlag string
)

func init() {
	flag.StringVar(&portFlag, "port", "8080", "Port to run the server on")
}
