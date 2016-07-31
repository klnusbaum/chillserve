package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"net/http"
)

func TestNewChillifierHandler(t *testing.T) {
	testMap := make(map[string]string)
	testMap["bad"] = "good"
	ch := NewChillifierHandler(testMap).(chillifierHandler)
	assert.Equal(t, "good", ch.replacer.Replace("bad"), "Good not replaced with bad")
	assert.Equal(t, "Cheese", ch.replacer.Replace("Cheese"), "Cheese was placed when it should not have been.")
}

func TestServeChillifierHttp(t *testing.T) {
	ch := chillifierHandler {
		replacer: strings.NewReplacer([]string{"bad", "good"}...),
	}

	req, _ := http.NewRequest("GET", "http://chill.com/chillify?text=bad", strings.NewReader(""))

	w := httptest.NewRecorder()
	ch.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, "{\"chill text\":\"good\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}

func TestServeChillifierHttpNoText(t *testing.T) {
	ch := chillifierHandler {
		replacer: strings.NewReplacer([]string{"bad", "good"}...),
	}

	req, _ := http.NewRequest("GET", "http://chill.com/chillify", strings.NewReader(""))

	w := httptest.NewRecorder()
	ch.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
	assert.Equal(t, "{\"error\":\"missing \\\"text\\\" parameter\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}
