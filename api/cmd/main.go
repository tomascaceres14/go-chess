package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/user"
)

func main() {

	sv := http.NewServeMux()

	sv.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	sv.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {

		var register user.RegisterUser
		if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
			JSONError(w, r, "Error decoding user", err, http.StatusBadRequest)
			return
		}

		if err := register.Validate(); err != nil {
			JSONError(w, r, "Error validating request", err, http.StatusBadRequest)
			return
		}

		response, err := json.Marshal(register)
		if err != nil {
			JSONError(w, r, "Error encoding user", err, http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	if err := http.ListenAndServe(":8080", sv); err != nil {
		log.Fatal(err)
	}
}

type HTTPError struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, r *http.Request, msg string, err error, status int) {
	log.Printf("ERROR %s: %s: %v", r.Pattern, msg, err)
	response, _ := json.Marshal(HTTPError{
		Error: msg,
	})
	w.WriteHeader(status)
	w.Write(response)
}
