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
	users, err := h.svc.GetAll(r.Context())
	if err != nil {
		utils.HTTPJsonError(w, r, "Error fetching users.", err, http.StatusInternalServerError)
	}
	utils.HTTPJsonResponse(w, users, http.StatusOK)
}
