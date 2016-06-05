package handlers

import (
	"net/http"
	"encoding/json"
	"strings"
)

type chillReplacer interface {
	Replace(s string) string
}

type chillifier struct {
	replacer chillReplacer
}

func NewChillifierHandler() http.Handler {
	return chillifier{
		replacer: strings.NewReplacer(
			" the ", " the chill ",
			" The ", " The chill ",
			"The ", "The chill ",
			" a ", " a chill ",
			" A ", " A chill ",
			"A ", "A chill ",
			" their ", " their chill ",
			" Their ", " Their chill ",
			" hate ", " love ",
			" Hate ", " Love ",
			"Hate ", "Love ",
			" my ", " my chill ",
			" My ", " My chill ",
			"My ", "My chill",
		),
	}
}

func (ch chillifier) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	ch.processRequest(rw, req)
}

func (ch chillifier) processRequest(rw http.ResponseWriter, req *http.Request) {
	text := req.FormValue("text")
	if text == "" {
		ch.writeErrorResponse(rw)
		return
	}

	ch.writeChillResponse(rw, text)
}

func (ch chillifier) writeErrorResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	response := map[string]string{"error": "missing \"text\" parameter"}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (ch chillifier) writeChillResponse(rw http.ResponseWriter, text string) {
	rw.WriteHeader(http.StatusOK)

	response := map[string]string{"chill text": ch.replacer.Replace(text)}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
