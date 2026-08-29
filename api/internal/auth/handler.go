package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/token"
	"github.com/tomascaceres14/go-chess/api/utils"
)

var (
	ErrUsernameTooShort   = errors.New("Username must be at least 6 characters long")
	ErrPasswordTooShort   = errors.New("Password must be at least 8 characters long")
	ErrPasswordsDontMatch = errors.New("Passwords do not match")
)

type Handler struct {
	svc           *Service
	tokenProvider token.TokenProvider
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var register UserRegister
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		utils.HTTPJsonError(w, r, "Error processing request", err, http.StatusBadRequest)
		return
	}

	user, err := h.svc.Register(r.Context(), register)
	if err != nil {
		switch {
		case errors.Is(err, ErrPasswordTooShort), errors.Is(err, ErrPasswordsDontMatch), errors.Is(err, ErrUsernameTooShort):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		default:
			utils.HTTPJsonError(w, r, "Internal server error", err, http.StatusInternalServerError)
		}
		return
	}

	token, err := h.tokenProvider.NewUserCredentials(user.ID)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error signing token", err, http.StatusBadRequest)
		return
	}

	utils.HTTPJsonResponse(w, token, http.StatusCreated)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials UserCredentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.HTTPJsonError(w, r, "Error processing request", err, http.StatusBadRequest)
		return
	}

	user, err := h.svc.userSvc.GetByUsername(r.Context(), credentials.Username)
	if err != nil {
		utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		return
	}

	token, err := h.tokenProvider.NewUserCredentials(user.ID)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error signing token", err, http.StatusBadRequest)
		return
	}

	utils.HTTPJsonResponse(w, token, http.StatusCreated)
}
