package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/utils"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	utils.HTTPJsonResponse(w, h.svc.GetAll(r.Context()), http.StatusOK)
}

func (h *Handler) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	var register RegisterUser
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

	response, err := json.Marshal(user)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error encoding json", err, http.StatusInternalServerError)
		return
	}

	utils.HTTPJsonResponse(w, response, http.StatusCreated)
}
