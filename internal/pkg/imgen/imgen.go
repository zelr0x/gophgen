// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import "image/color"

// Default values used by image generator.
// It is essentially an imgen control panel.
const (
	// DefaultGifColors defines the amount of colors used by a GIF generator.
	DefaultGifColors uint16 = 256
	// DefaultIsIdenticon defines if images should be identicons by default.
	DefaultIsIdenticon = false
	// MinIdenticonWidth defines minimum image width.
	MinIdenticonWidth int = 16
	// MaxWidth defines maximum image width.
	MaxWidth int = 4096
	// MaxHeight defines maximum image height.
	MaxHeight int = 2160
	// defaultWidth defines default image width.
	defaultWidth uint16 = 64
	// defaultHeight defines default image height.
	defaultHeight uint16 = 64
	// defaultDataSize defines data size in bytes.
	defaultDataSize uint16 = 128
	// defaultOpacity defines default opacity for image generator.
	defaultOpacity uint8 = 0xff
)

// defaultPalette defines default pallette to be used by other functions.
var defaultPalette = &x11

// Size represents a size of the image.
type Size struct {
	Width, Height uint16
}

// Color represents a color of the image.
type Color struct {
	BgColor, FgColor color.RGBA
}

// Data encapsulates underlying type of data generation seed.
type Data []byte

// Img represents image parameters.
type Img struct {
	Size  Size
	Color Color
	Ext   Extension
	Data  Data
}

// init initializes random generator. It is called on start automatically.
func init() {
	initRandom()
}

// DefaultSize returns Size with default values for width and height.
func DefaultSize() Size {
	return Size{defaultWidth, defaultHeight}
}

// DefaultColor returns Color with default values for BgColor and FgColor.
func DefaultColor() Color {
	return Color{RandomColor(), RandomPaleColor()}
}
