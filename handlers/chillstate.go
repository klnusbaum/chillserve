package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type stateChiller struct {
	imageUrls map[string]string
}

func NewStateChiller(stateImagesLocation string) http.Handler {
	imageUrls := make(map[string]string)
	for state := range validStates {
		imageUrls[state] = fmt.Sprintf("%s/%s.jpg", stateImagesLocation, state)
	}
	return stateChiller{
		imageUrls: imageUrls,
	}
}

func (sc stateChiller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	sc.processRequest(rw, req)
}

func (sc stateChiller) processRequest(rw http.ResponseWriter, req *http.Request) {
	state := req.FormValue("state")
	if state == "" {
		sc.writeErrorResponse(rw)
		return
	}

	if !validStates[state] {
		sc.writeInvalidStateResponse(rw, state)
	}

	sc.writeChillResponse(rw, state)
}

func (sc stateChiller) writeErrorResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	response := map[string]string{"error": "missing \"state\" parameter"}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (sc stateChiller) writeInvalidStateResponse(rw http.ResponseWriter, missingState string) {
	rw.WriteHeader(http.StatusNotFound)
	response := map[string]string{"error": fmt.Sprintf("Could not find state %q", missingState)}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (sc stateChiller) writeChillResponse(rw http.ResponseWriter, state string) {
	rw.WriteHeader(http.StatusOK)
	response := map[string]string{"chill image": sc.imageUrls[state]}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

