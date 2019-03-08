// Package server specifies available API routes,
// parses API requests and writes a response.
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	assert := assert.New(t)
	var methodTests = []struct {
		method           string
		expectedHttpCode int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusOK},
		{"TRACE", http.StatusOK},
		{"HEAD", http.StatusOK},
		{"PUT", http.StatusOK},
		{"DELETE", http.StatusOK},
		{"CONNECT", http.StatusOK},
		{"PATCH", http.StatusOK},
	}
	router := setupRouter()
	res := httptest.NewRecorder()
	for _, test := range methodTests {
		req, err := http.NewRequest(test.method, "/", nil)
		router.ServeHTTP(res, req)
		assert.Equal(nil, err, "An error occured upon receiving a request.")
		assert.Equal(test.expectedHttpCode, res.Code, "Wrong status code.")
	}
}
