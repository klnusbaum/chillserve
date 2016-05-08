package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type ChillHandler struct {
	ChillPhrases []string
}

func (ch ChillHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	writeChillHeader(rw)
	ch.writeChillJson(rw)
}

func writeChillHeader(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
}

func (ch ChillHandler) writeChillJson(rw http.ResponseWriter) {
	numOptions := len(ch.ChillPhrases)
	selectionIndex := rand.Intn(numOptions)
	response := map[string]string{"chill": ch.ChillPhrases[selectionIndex]}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
