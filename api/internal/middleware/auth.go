package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tomascaceres14/go-chess/api/utils"
)

var (
	ErrBadAuthHeaderFormat = errors.New("Invalid Authorization header format")
	ErrTokenInvalid        = errors.New("Error validating token")
)

func (m *Middleware) JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[AUTH CHECK] %s: In progress", r.Pattern)

		// Get Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check format
		s := strings.Split(authHeader, " ")
		if len(s) != 2 || s[0] != "Bearer" {
			utils.HTTPJsonError(w, r, ErrBadAuthHeaderFormat.Error(), ErrBadAuthHeaderFormat, http.StatusForbidden)
			return
		}

		// Validate token
		t, err := m.TokenProvider.ValidateToken(s[1])
		if err != nil {
			utils.HTTPJsonError(w, r, "Error validating token", err, http.StatusForbidden)
			return
		}

		// Get and validate ID
		id, err := t.Claims.GetSubject()
		if err != nil {
			utils.HTTPJsonError(w, r, "Invalid token", err, http.StatusForbidden)
			return
		}

		if !m.UserService.ExistsByID(r.Context(), id) {
			utils.HTTPJsonError(w, r, fmt.Sprintf("User ID: %s not found or banned", id), err, http.StatusForbidden)
			return
		}

		// Pass id in context
		ctx := context.WithValue(
			r.Context(),
			"userID",
			id,
		)

		log.Printf("[AUTH CHECK] %s: OK. UserID: %s", r.Pattern, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
