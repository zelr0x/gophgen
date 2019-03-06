// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMime(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("image/png", PNG.Mime(), "Wrong PNG Mime type.")
	assert.Equal("image/jpeg", JPEG.Mime(), "Wrong JPEG Mime type.")
	assert.Equal("image/gif", GIF.Mime(), "Wrong GIF Mime type.")
}

func TestGetExtensionNames(t *testing.T) {
	assert.Equal(t, []string{"png", "jpeg", "jpg", "gif"}, GetExtensionNames())
}

func TestGetExtension(t *testing.T) {
	defaultExt := PNG
	assert := assert.New(t)

	var tests = []struct {
		input          string
		expectedExt    Extension
		expectedExists bool
	}{
		{"\000", defaultExt, false},
		{"\n", defaultExt, false},
		{"\r", defaultExt, false},
		{"\n\r", defaultExt, false},
		{"\r\n", defaultExt, false},
		{"嘊", defaultExt, false},
		{"", defaultExt, false},
		{"a", defaultExt, false},
		{"asmdo8j120ed9kad", defaultExt, false},
		{"asm do8j 120e d9kad .. .", defaultExt, false},
		{"-1000000000000", defaultExt, false},
		{"-32768", defaultExt, false},
		{"-1", defaultExt, false},
		{"0", defaultExt, false},
		{"1024", defaultExt, false},
		{"1025", defaultExt, false},
		{"8080", defaultExt, false},
		{"65535", defaultExt, false},
		{"65636", defaultExt, false},
		{"1000000000000", defaultExt, false},
		{"png", defaultExt, true},
		{"PNG", defaultExt, true},
		{"jpegG", defaultExt, false},
		{"jpeg", JPEG, true},
		{"jpg", JPEG, true},
		{"jpG", JPEG, true},
		{"gif", GIF, true},
	}

	for _, test := range tests {
		var ext Extension
		ext, exists := GetExtension(test.input)
		assert.Equal(test.expectedExt, ext, "Wrong extension.")
		if exists {
			assert.Equal(test.expectedExists, exists, "Nonexistent extension was found.")
		} else {
			assert.Equal(test.expectedExists, exists, "Existent extension was not found.")
		}
	}
}
