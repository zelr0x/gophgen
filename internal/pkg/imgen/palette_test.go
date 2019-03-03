// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColor(t *testing.T) {
	assert := assert.New(t)
	testPalette := &x11
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
		color, exists := testPalette.getColor(test.inputColorName)
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

func TestRandom(t *testing.T) {
	testPalette := &x11
	testCases := 10
	tests := make([]color.RGBA, 0, testCases)
	for testCase := 0; testCase < testCases; testCase++ {
		tests = append(tests, testPalette.random())
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
