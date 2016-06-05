package handlers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

func TestNewRandomChillHandler(t *testing.T) {
	ch := NewRandomChillHandler("phrase1", "phrase2").(randomChillHandler)
	assert.Len(t, ch.phrases, 2, "Phrases not of length 2")
	assert.Equal(t, []string{"phrase1", "phrase2"}, ch.phrases, "Phrase list unequal")
}

func TestServeRandomHttp(t *testing.T) {
	ch := randomChillHandler{
		phrases:[]string{"phrase1", "phrase2"},
		random:func(n int) int { return 0 },
	}

	w := httptest.NewRecorder()
	ch.ServeHTTP(w, nil)

	assert.Equal(t, "{\"chill\":\"phrase1\"}\n", w.Body.String(), "Incorrect response")
	assert.Len(t,  w.HeaderMap["Content-Type"], 1, "Multiple content types")
	assert.Equal(t, "text/json; charset=utf-8", w.HeaderMap["Content-Type"][0], "Incorrect content type")
}
