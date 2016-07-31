package handlers

import (
	"net/http"
	"encoding/json"
	"strings"
)

type chillifierHandler struct {
	replacer *strings.Replacer
}

func NewChillifierHandler(replacements map[string]string) http.Handler {
	replacerArgs := make([]string, len(replacements) * 2)
	for k, v := range replacements {
		replacerArgs = append(replacerArgs, k, v)
	}
	return chillifierHandler{
		replacer: strings.NewReplacer(replacerArgs...),
	}
}

func (ch chillifierHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	ch.processRequest(rw, req)
}

func (ch chillifierHandler) processRequest(rw http.ResponseWriter, req *http.Request) {
	text := req.FormValue("text")
	if text == "" {
		ch.writeBadRequestResponse(rw)
		return
	}

	ch.writeChillResponse(rw, text)
}

func (ch chillifierHandler) writeBadRequestResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	response := map[string]string{"error": "missing \"text\" parameter"}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (ch chillifierHandler) writeChillResponse(rw http.ResponseWriter, text string) {
	rw.WriteHeader(http.StatusOK)

	response := map[string]string{"chill text": ch.replacer.Replace(text)}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
