// Package server specifies available API routes, concurrently parses
// API requests and writes a response using gin-gonic.
package server

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zelr0x/gophgen/internal/pkg/imgen"
)

// parseAndWrite parses a url and creates an imgen.Img
// which is then passed to write().
func parseAndWrite(c *gin.Context, first string, rest string) {
	// In the API, identicon flag is optional but since it comes first,
	// offset for all subsequent values should be handled appropriately.
	// We need to know if an image is identicon in advance
	// hence parser.IsIdenticon is called without a goroutine.
	isIdenticon := isIdenticon(first[0])
	offset := 0
	if isIdenticon {
		offset = 1
	}

	img := imgen.Img{}
	sizeCh := make(chan imgen.Size)

	if len(rest) != 0 {
		p := strings.Split(rest, "/") // p[0] is a space.
		if isIdenticon {
			go parseSize(p[1], isIdenticon, sizeCh)
		} else {
			go parseSize(first, isIdenticon, sizeCh)
		}

		bgColCh := make(chan color.RGBA)
		fgColCh := make(chan color.RGBA)
		extCh := make(chan imgen.Extension)
		dataCh := make(chan imgen.Data)

		plen := len(p)
		if plen > 1+offset {
			go parseColor(p[1+offset], bgColCh)
			if plen > 2+offset {
				go parseColor(p[2+offset], fgColCh)
				img.Color.FgColor = <-fgColCh
			} else {
				img.Color.FgColor = imgen.RandomPaleColor()
			}
			img.Color.BgColor = <-bgColCh
		} else {
			img.Color = imgen.DefaultColor()
		}
		go parseExt(p[1+offset:], extCh)
		go parseData(p[plen-1], isIdenticon, dataCh)

		img.Size = <-sizeCh
		img.Ext = <-extCh
		img.Data = <-dataCh
	} else {
		// Default value for ext is correct.
		img.Size = imgen.DefaultSize()
		img.Color = imgen.DefaultColor()
		img.Data = imgen.DefaultData(isIdenticon)
	}
	write(c, &img, isIdenticon)
}

// write creates an image.Image from imgen.Img,
// encodes it and writes it as a file to gin.context.
func write(c *gin.Context, img *imgen.Img, isIdenticon bool) {
	var image image.Image
	if isIdenticon {
		image = imgen.NewIdenticon(img)
	} else {
		image = imgen.NewImage(img)
	}

	buffer := new(bytes.Buffer)
	err := encodeToBuffer(image, img.Ext, buffer)
	if err != nil {
		log.Println("Unable to encode image")
	}

	// To use bytes in html img tag, they should be sent as an attachment.
	c.Header("Content-Disposition", "attachment; filename=\"image."+
		img.Ext.String()+"\"")
	c.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	c.Data(http.StatusOK, img.Ext.Mime(), buffer.Bytes())
}

// encodeToBuffer encodes image using appropriate extension codec and
// writes resulting bytes to a buffer passed by reference.
func encodeToBuffer(image image.Image, ext imgen.Extension, buffer *bytes.Buffer) error {
	var err error
	switch ext {
	case imgen.JPEG:
		err = jpeg.Encode(buffer, image, nil)
	case imgen.GIF:
		options := gif.Options{}
		options.NumColors = int(imgen.DefaultGifColors)
		err = gif.Encode(buffer, image, &options)
	default:
		err = png.Encode(buffer, image)
	}
	return err
}
