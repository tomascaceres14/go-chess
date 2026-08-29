package middleware

import (
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/auth"
	"github.com/tomascaceres14/go-chess/api/internal/user"
)

type Middleware struct {
	TokenProvider auth.TokenProvider
	UserService   *user.Service
}

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func (m *Middleware) Use(h http.HandlerFunc, middlewares ...MiddlewareFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
