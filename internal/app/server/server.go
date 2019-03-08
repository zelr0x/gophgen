// Package server specifies available API routes,
// parses API requests and writes a response.
package server

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// setupRouter specifies available API routes.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(os.Getenv("GOPATH"),
			"src/github.com/zelr0x/gophgen/web/static/index.html"))
	})
	r.GET("/:first", func(c *gin.Context) {
		parseAndWrite(c, c.Param("first"), "")
	})
	r.GET("/:first/*rest", func(c *gin.Context) {
		parseAndWrite(c, c.Param("first"), c.Param("rest"))
	})

	return r
}

// Serve starts the router.
func Serve(port uint16) {
	r := setupRouter()
	r.Run(":" + strconv.Itoa(int(port)))
}
