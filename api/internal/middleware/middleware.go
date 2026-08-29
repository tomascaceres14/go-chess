package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/utils"
)

func Use(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		for _, v := range m {
			if err = v(w, r); err != nil {
				break
			}
		}

		if err != nil {
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
			return
		}

		h(w, r)
	}

}

type Middleware func(w http.ResponseWriter, r *http.Request) error

func Mid1(w http.ResponseWriter, r *http.Request) error {
	log.Printf("Mid1 running")
	return nil
}

func Mid2(w http.ResponseWriter, r *http.Request) error {
	log.Println("Mid2 running")
	log.Println("Obtaining bearer token")
	token := r.Header.Get("Authorization")
	
	if token == "" {
		return errors.New("No Authorization token provided")
	}
	return nil
}

func Mid3(w http.ResponseWriter, r *http.Request) error {
	log.Println("Mid3 running")
	fmt.Println("mid1 and 2 successful. Adding records and proceding")
	return nil
}
