package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type HTTPError struct {
	Error string `json:"error"`
}

func HTTPJsonError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Printf("--- ERROR %s | Status: %d |Err: %v", r.Pattern, status, err)

	HTTPJsonResponse(
		w,
		HTTPError{
			Error: msg,
		},
		status,
	)

}

func HTTPJsonResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
