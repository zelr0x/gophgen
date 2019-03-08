// Package server specifies available API routes,
// parses API requests and writes a response.
package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zelr0x/gophgen/internal/pkg/imgen"
)

func TestIsIdenticon(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    byte
		expected bool
	}{
		{'\000', false},
		{'\n', false},
		{'\r', false},
		{' ', false},
		{'a', false},
		{'J', false},
		{'0', false},
		{'i', true},
		{'I', true},
	}

	for _, test := range tests {
		assert.Equal(test.expected, isIdenticon(test.input),
			"Identicon specifier was not identified properly.")
	}
}

func TestParseSize(t *testing.T) {
	assert := assert.New(t)
	defaultSize := imgen.DefaultSize()

	var tests = []struct {
		input        string
		expectedSize imgen.Size
	}{
		{"\000", defaultSize},
		{"\n", defaultSize},
		{"\r", defaultSize},
		{"\n\r", defaultSize},
		{"\r\n", defaultSize},
		{"嘊", defaultSize},
		{"", defaultSize},
		{"a", defaultSize},
		{"-10000000 00000", defaultSize},
		{"-32768", defaultSize},
		{"0", defaultSize},
		{"8080", defaultSize},
		{"65535", defaultSize},
		{"65636", defaultSize},
		{"1000000000000", defaultSize},
		{"1024", imgen.Size{1024, 1024}},
	}

	ch := make(chan imgen.Size)

	for _, test := range tests {
		go parseSize(test.input, false, ch)
		assert.Equal(test.expectedSize, <-ch, "Wrong image size for input: "+test.input+".")
		go parseSize(test.input, true, ch)
		assert.Equal(test.expectedSize, <-ch, "Wrong identicon size for input: "+test.input+".")
	}

	var imageOnlyTests = []struct {
		input        string
		expectedSize imgen.Size
	}{
		{"-1", imgen.Size{1, 1}},
		{"10", imgen.Size{10, 10}},
		{"10x10", imgen.Size{10, 10}},
	}

	for _, test := range imageOnlyTests {
		go parseSize(test.input, false, ch)
		assert.Equal(test.expectedSize, <-ch, "Wrong image size for input: "+test.input+".")
	}

	var identiconOnlyTests = []struct {
		input        string
		expectedSize imgen.Size
	}{
		{"-1", defaultSize},
		{"10", defaultSize},
		{"10x10", defaultSize},
	}

	for _, test := range identiconOnlyTests {
		go parseSize(test.input, true, ch)
		assert.Equal(test.expectedSize, <-ch, "Wrong identicon size for input: "+test.input+".")
	}
}
