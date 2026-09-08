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
	ErrWrongPassword      = errors.New("Wrong password")
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

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var register UserRegister
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		utils.HTTPJsonError(w, r, "Error processing request", err, http.StatusBadRequest)
		return
	}

	creds, err := h.svc.Register(r.Context(), register)
	if err != nil {
		switch {
		case errors.Is(err, ErrPasswordTooShort), errors.Is(err, ErrPasswordsDontMatch), errors.Is(err, ErrUsernameTooShort):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		default:
			utils.HTTPJsonError(w, r, "Internal server error", err, http.StatusInternalServerError)
		}
		return
	}

	utils.HTTPJsonResponse(w, creds, http.StatusCreated)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var login UserLogin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.HTTPJsonError(w, r, "Error processing request", err, http.StatusBadRequest)
		return
	}

	creds, err := h.svc.Login(r.Context(), login)
	if err != nil {
		utils.HTTPJsonError(w, r, err.Error(), err, http.StatusUnauthorized)
		return
	}

	utils.HTTPJsonResponse(w, creds, http.StatusCreated)
}

func (h *Handler) HandleRefresh(w http.ResponseWriter, r *http.Request) {

}
