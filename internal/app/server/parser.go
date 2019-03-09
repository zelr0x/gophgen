// Package server specifies available API routes,
// parses API requests and writes a response.
package server

import (
	"image/color"
	"regexp"
	"strconv"
	"strings"

	"github.com/zelr0x/gophgen/internal/pkg/imgen"
)

// isIdenticon returns true if given byte is an identicon specifier and false otherwise.
func isIdenticon(b byte) bool {
	for _, specifier := range []byte{'i', 'I'} {
		if specifier == b {
			return true
		}
	}
	return false
}

// parseSize tries to extract main width and height from a given string.
// Assigns default values on fail.
func parseSize(s string, isIdenticon bool, sizeCh chan imgen.Size) {
	re := regexp.MustCompile(`\d+`)
	dimensions := re.FindAllString(s, -1)
	dimensionsLen := len(dimensions)

	if dimensionsLen < 1 {
		sizeCh <- imgen.DefaultSize()
		return
	}

	width := 0
	height := 0
	if parsed, err := strconv.Atoi(dimensions[0]); err == nil {
		width = parsed
	}

	if dimensionsLen == 1 {
		height = width
	} else if parsed, err := strconv.Atoi(dimensions[1]); err == nil {
		height = parsed
	}

	size := imgen.Size{}
	if width > imgen.MaxWidth || width < 1 ||
		height > imgen.MaxHeight || height < 1 ||
		(width < imgen.MinIdenticonWidth && isIdenticon) {
		sizeCh <- imgen.DefaultSize()
		return
	}
	size.Width = uint16(width)
	size.Height = uint16(height)
	sizeCh <- size
}

// parseColor tries to extract color.RGBA from a given string,
// returns random color.RGBA on fail.
func parseColor(s string, colCh chan color.RGBA) {
	rgba, exists := imgen.GetColor(s)
	if exists {
		colCh <- rgba
		return
	}

	re := regexp.MustCompile("[[:xdigit:]]")
	matches := re.FindAllString(s, -1)
	if matches != nil {
		colCh <- imgen.Hex2rgba(strings.Join(matches, ""))
		return
	}

	colCh <- imgen.RandomColor()
}

// parseExt tries to extract imgen.Extension from a string s.
func parseExt(s []string, extCh chan imgen.Extension) {
	var ext imgen.Extension // Correct default value is provided by iota.
	for _, str := range s {
		k, exists := imgen.GetExtension(str)
		if exists {
			ext = k
		}
	}
	extCh <- ext
}

// parseData tries to extract data []byte from a string s, returns imgen.DefData on fail.
func parseData(s string, isIdenticon bool, dataCh chan imgen.Data) {
	for _, specifier := range []string{"d=", "D="} {
		dindex := strings.LastIndex(s, specifier)
		if dindex != -1 && dindex != len(s) {
			// "=" is used as a part of the seed for identicon generator
			// but it is not a problem since it's not a crypto app.
			dataCh <- []byte(s[dindex:])
		}
	}
	dataCh <- imgen.DefaultData(isIdenticon)
}
