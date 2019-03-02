// Package server specifies available API routes,
// parses API requests and writes a response.
package server

import "fmt"

func Serve(portNumber uint16) {
	port := ":" + string(portNumber)
	fmt.Println(port)
}
