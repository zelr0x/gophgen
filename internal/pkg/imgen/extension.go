// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import "strings"

// Extension is an enumerator of supported image extensions.
type Extension uint8

const (
	// PNG is a default image extension in imgen.
	PNG Extension = iota
	// JPEG is another image extension supported by imgen.
	JPEG
	// GIF is another image extension supported by imgen.
	GIF
)

// Mime returns a respective MIME-type string for a given Extension.
func (e Extension) Mime() string {
	main := "image/"
	switch e {
	case JPEG:
		return main + "jpeg"
	case GIF:
		return main + "gif"
	default:
		return main + "png"
	}
}

// GetExtensionNames return a []string with available extensions.
func GetExtensionNames() []string {
	return []string{"png", "jpeg", "jpg", "gif"}
}

var extensions = map[string]Extension{
	"png":  PNG,
	"jpeg": JPEG,
	"jpg":  JPEG,
	"gif":  GIF,
}

// GetExtension tries to extract Extension from s.
// Returns true on success, false otherwise.
func GetExtension(s string) (Extension, bool) {
	ext, exists := extensions[strings.ToLower(s)]
	return ext, exists
}
