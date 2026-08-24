package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type HTTPError struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Printf("--- ERROR %s | Status: %d |Err: %v", r.Pattern, status, err)

	JSONResponse(
		w,
		HTTPError{
			Error: msg,
		},
		status,
	)

}

func JSONResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
