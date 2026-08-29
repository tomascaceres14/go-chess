package user

import (
	"errors"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/token"
	"github.com/tomascaceres14/go-chess/api/utils"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

type Handler struct {
	svc           *Service
	tokenProvider token.TokenProvider
}

func NewHandler(svc *Service, tokenProvider token.TokenProvider) *Handler {
	return &Handler{
		svc:           svc,
		tokenProvider: tokenProvider,
	}
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	utils.HTTPJsonResponse(w, h.svc.GetAll(r.Context()), http.StatusOK)
}
