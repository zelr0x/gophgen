// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"image"

	"github.com/issue9/identicon"
)

// NewIdenticon creates a new image.Image identicon with given parameters.
func NewIdenticon(i *Img) image.Image {
	img, err := identicon.Make(int(i.Size.Width), i.Color.BgColor, i.Color.FgColor, i.Data)
	if err != nil {
		panic("Possible memory problem detected.")
	}
	return img
}
