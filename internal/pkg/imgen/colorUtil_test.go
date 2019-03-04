// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHex2rgba tests only valid inputs because all others are discarded
// and in such case random color is returned.
func TestHex2rgba(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    string
		expected color.RGBA
	}{
		{"ffffffff", color.RGBA{0xff, 0xff, 0xff, 0xff}},
		{"abcdef55", color.RGBA{0xab, 0xcd, 0xef, 0x55}},
		{"15a3ef", color.RGBA{0x15, 0xa3, 0xef, 0xff}},
		{"ffff", color.RGBA{0xff, 0xff, 0xff, 0xff}},
		{"efe", color.RGBA{0xee, 0xff, 0xee, 0xff}},
		{"010", color.RGBA{0x00, 0x11, 0x00, 0xff}},
	}

	for _, test := range tests {
		assert.Equal(test.expected, Hex2rgba(test.input), "Wrong color detected.")
	}
}

func TestGetColor(t *testing.T) {
	assert := assert.New(t)
	nilRGBA := color.RGBA{0, 0, 0, 0}

	var tests = []struct {
		inputColorName string
		expectedColor  color.RGBA
		expectedExists bool
	}{
		{"\000", nilRGBA, false},
		{"\n", nilRGBA, false},
		{"\r", nilRGBA, false},
		{"\n\r", nilRGBA, false},
		{"\r\n", nilRGBA, false},
		{"嘊", nilRGBA, false},
		{"", nilRGBA, false},
		{"a", nilRGBA, false},
		{"asmdo8j120ed9kad", nilRGBA, false},
		{"asm do8j 120e d9kad .. .", nilRGBA, false},
		{"-1000000000000", nilRGBA, false},
		{"-1", nilRGBA, false},
		{"0", nilRGBA, false},
		{"1000000000000", nilRGBA, false},
		{"red", color.RGBA{0xff, 0x00, 0x00, 0xff}, true},
	}

	for _, test := range tests {
		color, exists := GetColor(test.inputColorName)
		assert.Equal(test.expectedColor, color, "Colors don't match")
		var msg string
		if exists {
			msg = "Returned color that is not present in the palette."
		} else {
			msg = "Didn't find color that is present in the palette."
		}
		assert.Equal(test.expectedExists, exists, msg)
	}
}

func TestRandomColor(t *testing.T) {
	testCases := 10
	tests := make([]color.RGBA, 0, testCases)
	for testCase := 0; testCase < testCases; testCase++ {
		tests = append(tests, RandomColor())
	}

	collisions := 0
	prevColor := color.RGBA{0, 0, 0, 0}
	for _, color := range tests {
		if color == prevColor {
			collisions++
		}
		prevColor = color
	}
	assert.True(t, collisions < (testCases/5), "Not enough random.")
}

// When PaleColor will be implemented properly, pale-check should be added.
func TestRandomPaleColor(t *testing.T) {
	testCases := 10
	tests := make([]color.RGBA, 0, testCases)
	for testCase := 0; testCase < testCases; testCase++ {
		tests = append(tests, RandomPaleColor())
	}

	collisions := 0
	prevColor := color.RGBA{0, 0, 0, 0}
	for _, color := range tests {
		if color == prevColor {
			collisions++
		}
		prevColor = color
	}
	assert.True(t, collisions < (testCases/5), "Not enough random.")
}
