package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

type stateChillHandler struct {
	imageUrls map[string]string
}

// NewStateChillHandler creates an http handler that will return chill images of states. The state image returned is
// determined by "state" parameter provided in the request.
//
// Example request:
// http://localhost:8008/states_chill?state=MN
//
// Example response:
//   {"chill image":"http://blahblahblh.cloudfront.net/CA.jpg"}
func NewStateChillHandler(stateImagesLocation string) http.Handler {
	imageUrls := make(map[string]string)
	for state := range validStates {
		imageUrls[state] = fmt.Sprintf("%s/%s.jpg", stateImagesLocation, state)
	}
	return stateChillHandler{
		imageUrls: imageUrls,
	}
}

func (sc stateChillHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/json; charset=utf-8")
	sc.processRequest(rw, req)
}

func (sc stateChillHandler) processRequest(rw http.ResponseWriter, req *http.Request) {
	state := strings.ToUpper(req.FormValue("state"))
	if state == "" {
		sc.writeBadRequestResponse(rw)
		return
	}

	if !validStates[state] {
		sc.writeInvalidStateResponse(rw, state)
	}

	sc.writeChillResponse(rw, state)
}

func (sc stateChillHandler) writeBadRequestResponse(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusBadRequest)
	response := map[string]string{"error": "missing \"state\" parameter"}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (sc stateChillHandler) writeInvalidStateResponse(rw http.ResponseWriter, missingState string) {
	rw.WriteHeader(http.StatusNotFound)
	response := map[string]string{"error": fmt.Sprintf("Could not find state %q", missingState)}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func (sc stateChillHandler) writeChillResponse(rw http.ResponseWriter, state string) {
	rw.WriteHeader(http.StatusOK)
	response := map[string]string{"chill image": sc.imageUrls[state]}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

