// Package main parses command-line arguments and starts the server.
package main

import (
	"os"
	"strconv"

	"github.com/zelr0x/gophgen/internal/app/server"
)

const (
	// defaultPort defines a port to use when none is specified on start.
	defaultPort uint16 = 8080
	// lastReserved is the highest reserved port number.
	lastReserved = 1024
)

// main starts the server on port.
func main() {
	server.Serve(getPort())
}

// getPort parses command-line arguments and returns port number to use.
func getPort() uint16 {
	if len(os.Args) > 1 {
		if p, err := strconv.ParseUint(os.Args[1], 10, 16); err == nil && p > lastReserved {
			return uint16(p)
		}
	}
	return defaultPort
}
