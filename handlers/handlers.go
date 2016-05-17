package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// A function that generates a random number between 0 and a given number
type randomer func(int) int

type chillHandler struct {
	phrases []string
	random randomer
}

func NewChillHandler(phrases ...string) http.Handler {
	return chillHandler{
		phrases:phrases,
		random:rand.Intn,
	}
}

func (ch chillHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	writeHeader(rw)
	ch.writeBody(rw)
}

func writeHeader(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
}

func (ch chillHandler) writeBody(rw http.ResponseWriter) {
	numOptions := len(ch.phrases)
	selectionIndex := ch.random(numOptions)
	response := map[string]string{"chill": ch.phrases[selectionIndex]}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
