package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"fmt"
	"strings"
	"net/http"
)

func TestNewStateChillHandler(t *testing.T) {
	sch := NewStateChillHandler("http://zombo.com").(stateChillHandler)
	assert.Len(t, sch.imageUrls, 50, "Image URLs not of length 50.")
	for state, imageLocation := range sch.imageUrls {
		assert.Equal(t, fmt.Sprintf("%s/%s.jpg", "http://zombo.com", state), imageLocation)
	}
}

func TestServeStateHttp(t *testing.T) {
	sch := makeTestStateChillHandler()

	req, _ := http.NewRequest("GET", "http://chill.com/states_chill?state=CA", strings.NewReader(""))
	w := httptest.NewRecorder()
	sch.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, "{\"chill_image\":\"http://zombo.com/CA.jpg\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}

func TestServeStateHttpNoState(t *testing.T) {
	sch := makeTestStateChillHandler()

	req, _ := http.NewRequest("GET", "http://chill.com/states_chill", strings.NewReader(""))
	w := httptest.NewRecorder()
	sch.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, "{\"error\":\"missing \\\"state\\\" parameter\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}

func TestServeStateHttpMissingState(t *testing.T) {
	sch := makeTestStateChillHandler()

	req, _ := http.NewRequest("GET", "http://chill.com/states_chill?state=PO", strings.NewReader(""))
	w := httptest.NewRecorder()
	sch.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 404)
	assert.Equal(t, "{\"error\":\"Could not find state \\\"PO\\\"\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}

func makeTestStateChillHandler() stateChillHandler {
	return stateChillHandler {
		imageUrls: map[string]string{
			"CA": "http://zombo.com/CA.jpg",
		},
	}
}
