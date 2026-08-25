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
	var params NewMatchParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.HTTPJsonError(w, r, "Error decoding body", err, http.StatusBadRequest)
		return
	}

	if params.UserID == "" {
		utils.HTTPJsonError(w, r, "User ID not found", nil, http.StatusBadRequest)
		return
	}

	match, err := h.svc.StartNewMatch(r.Context(), params.UserID, params.Whites)
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
func (h *Handler) HandleGameWebSocket(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("gameID")
	userID := r.URL.Query().Get("userID")

	m, err := h.svc.AssignPlayerToMatch(r.Context(), gameID, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrMatchNotFound):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		default:
			utils.HTTPJsonError(w, r, "Internal server error", err, http.StatusInternalServerError)
		}
		return
	}

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		conn.WriteJSON(GameResponse{
			Command: MatchWaitingStatus,
			Valid:   true,
		})

		go func() {
			for msg := range m.ResponsesCh {
				if err := conn.WriteJSON(msg); err != nil {
					log.Println(err)
					return
				}
			}
		}()

		var command GameCommand
		if err := conn.ReadJSON(&command); err != nil {
			log.Println(err)
		}
		m.CommandsCh <- command
	}
}
