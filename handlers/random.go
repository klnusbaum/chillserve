package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// A function that generates a random number between 0 and a given number
type randomer func(int) int

type randomChillHandler struct {
	phrases []string
	random randomer
}

// NewChillHandler creates an http handler that will give simple, chill responses. Responses are in the form
// of a json map. There is one key, "chill", and it's value will be a chill phrase. The chill phrase will be choosen
// at random from one of the phrases provided.
//
// Example response:
//   { "chill" : "super chill" }
func NewRandomChillHandler(phrases ...string) http.Handler {
	return randomChillHandler{
		phrases:phrases,
		random:rand.Intn,
	}
}

func (rch randomChillHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	rch.writeBody(rw)
}

func (rch randomChillHandler) writeBody(rw http.ResponseWriter) {
	numOptions := len(rch.phrases)
	selectionIndex := rch.random(numOptions)
	response := map[string]string{"chill": rch.phrases[selectionIndex]}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
