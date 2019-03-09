// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"image/color"
)

// Hex2rgba converts string containing hex color code to respective color.RGBA.
// Expects length of 3 to 8 bytes without leading hash sign.
func Hex2rgba(s string) color.RGBA {
	col := color.RGBA{}
	slen := len(s)
	if slen == 8 || slen == 6 {
		col.R = hex2uint8(s[:2])
		col.G = hex2uint8(s[2:4])
		col.B = hex2uint8(s[4:6])
		if slen == 8 {
			col.A = hex2uint8(s[6:])
		} else {
			col.A = defaultOpacity
		}
	} else if slen == 4 || slen == 3 {
		col.R = hex2uint8(twice(s[0]))
		col.G = hex2uint8(twice(s[1]))
		col.B = hex2uint8(twice(s[2]))
		if slen == 4 {
			col.A = hex2uint8(twice(s[3]))
		} else {
			col.A = defaultOpacity
		}
	} else {
		col = RandomColor()
	}
	return col
}

// GetColor returns color.Color for a given color name
// and true if color was found or false otherwise.
// It's purpose is to encapsulate pallettes inside imgen.
func GetColor(name string) (color.RGBA, bool) {
	color, exists := defaultPalette.getColor(name)
	return color, exists
}

// RandomColor returns random color.RGBA.
func RandomColor() color.RGBA {
	return defaultPalette.random()
}

// RandomPaleColor returns random pale color.RGBA.
// TODO: Implement this feature.
func RandomPaleColor() color.RGBA {
	return defaultPalette.random()
}
