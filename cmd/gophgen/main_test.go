// Package main parses command-line arguments and starts the server.
package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort(t *testing.T) {
	var defaultPort uint16 = 8080
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected uint16
	}{
		{"\000", defaultPort},
		{"\n", defaultPort},
		{"", defaultPort},
		{"a", defaultPort},
		{"asmdo8j120ed9kad", defaultPort},
		{"-1000000000000", defaultPort},
		{"-32768", defaultPort},
		{"-1", defaultPort},
		{"0", defaultPort},
		{"1024", defaultPort},
		{"1025", 1025},
		{"8080", defaultPort},
		{"65535", 65535},
		{"65636", defaultPort},
		{"1000000000000", defaultPort},
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	for _, test := range tests {
		os.Args = []string{"cmd", test.input}
		assert.Equal(getPort(), test.expected)
	}
}
