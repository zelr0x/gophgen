// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"image"
	"image/draw"
)

// NewImage creates a new image.Image with given parameters.
func NewImage(i *Img) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, int(i.Size.Width), int(i.Size.Height)))
	draw.Draw(img, img.Bounds(), &image.Uniform{i.Color.BgColor}, image.ZP, draw.Src)
	return img
}
