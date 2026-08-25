package match

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/websocket"
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

func (h *Handler) HandleNewMatch(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	colour := r.URL.Query().Get("colour")

	if userID == "" {
		utils.HTTPJsonError(w, r, "User ID not found", nil, http.StatusBadRequest)
		return
	}

	if colour == "" {
		utils.HTTPJsonError(w, r, "Must provide chosen color: 1 for Whites, 0 for Blacks", nil, http.StatusBadRequest)
		return
	}

	var whites bool
	switch colour {
	case "1":
		whites = true
	case "0":
		whites = false
	default:
		utils.HTTPJsonError(w, r, "Incorrect colour option: 1 for Whites, 0 for Blacks", nil, http.StatusBadRequest)
		return
	}

	match, err := h.svc.StartNewMatch(r.Context(), userID, whites)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error creating match", err, http.StatusInternalServerError)
		return
	}

	wsPath := fmt.Sprintf("/ws/game/%s", match.ID)

	json.NewEncoder(w).Encode(
		map[string]any{
			"ws_path": wsPath,
			"token":   "token123_owner123",
		},
	)
}

// localhost:80/ws/game/{gameID} Bearer: Authorization token
func (h *Handler) HandleGameWebSocket(w http.ResponseWriter, r *http.Request, manager *websocket.ConnectionManager) {
	gameID := r.PathValue("gameID")
	userID := r.URL.Query().Get("userID")

	match, err := h.svc.AssignPlayerToMatch(r.Context(), gameID, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrMatchNotFound):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		default:
			utils.HTTPJsonError(w, r, "Internal server error", err, http.StatusInternalServerError)
		}
		return
	}

	conn, err := manager.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		if match.Status != StatusOngoing {
			continue
		}

		var command GameCommand
		if err := conn.ReadJSON(&command); err != nil {
			log.Println(err)
			return
		}

		command.UserID = userID
	}
}
