// Package handlers wires HTTP requests to the greeter package.
package handlers

import (
	"encoding/json"
	"net/http"

	"greeter/internal/greeter"
)

type greetingResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// Hello handles GET /hello.
func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	message, err := greeter.Greet(name)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, greetingResponse{Message: message})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
