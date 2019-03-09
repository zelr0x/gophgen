// Package imgen defines constants used by image generator
// as well as methods of creating images with different parameters.
// It also contains color palettes.
package imgen

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// initRandom initializes random generator for color picking function.
func initRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// DefaultData returns byte array of size suited for an image type.
// It is used to generate random data when none is provided by the user.
func DefaultData(isIdenticon bool) []byte {
	if isIdenticon {
		return randomBytes(int(defaultDataSize))
	}
	return randomBytes(int(defaultDataSize))
}

// hex2uint8 converts given string containing hexadecimal integer to uint8.
func hex2uint8(s string) uint8 {
	i, err := strconv.ParseUint(s, 16, 8)
	if err != nil {
		return uint8(rand.Intn(255))
	}
	return uint8(i)
}

// twice returns string containing given byte b twice.
func twice(b byte) string {
	return strings.Repeat(string(b), 2)
}

// randomBytes returns random byte array of given length.
func randomBytes(dataSize int) []byte {
	data := make([]byte, dataSize)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}
