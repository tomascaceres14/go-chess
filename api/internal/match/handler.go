package match

import (
	"encoding/json"
	"errors"
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

func (h *Handler) HandleGameWebSocket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	matchID := r.PathValue("matchID")
	userID := ctx.Value("userID").(string)

	if userID == "" {
		utils.HTTPJsonError(w, r, "User ID not found", nil, http.StatusBadRequest)
		return
	}

	responseCh, err := h.svc.AddUserToMatch(ctx, matchID, userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrMatchNotFound), errors.Is(err, ErrOwnerNotConnected):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		case errors.Is(err, ErrUserAlreadyConnected):
			utils.HTTPJsonError(w, r, err.Error(), err, http.StatusConflict)
		default:
			utils.HTTPJsonError(w, r, "Internal server error", err, http.StatusInternalServerError)
		}
		return
	}
	defer h.svc.RemoveUserFromMatch(matchID, userID)

	commandsCh, err := h.svc.GetCommandsCh(matchID, userID)
	if err != nil {
		utils.HTTPJsonError(w, r, err.Error(), err, http.StatusBadRequest)
		return
	}

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	conn.WriteJSON(GameResponse{
		Command: MatchWaitingStatus,
		Valid:   true,
	})

	// Redirects incoming match messages to user
	go func() {
		for msg := range responseCh {

			if err := conn.WriteJSON(msg); err != nil {
				log.Printf("Error writing WS message: %v", err)
				return
			}

			if msg.Command == MatchEndCmd {
				log.Printf("Match ID: %s finalized.", matchID)
				h.svc.FinalizeMatch(ctx, matchID, msg.Status, msg.Data["FEN"].(string), msg.Grid)
				return
			}
		}
	}()

	// Reads incoming user messages and redirects to match
	for {
		var command GameCommand
		if err := conn.ReadJSON(&command); err != nil {
			log.Printf("Error reading WS message: %v", err)
			return
		}

		command.UserID = userID
		commandsCh <- command
	}
}

func (h *Handler) HandleNewMatch(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		utils.HTTPJsonError(w, r, "User ID not found", nil, http.StatusBadRequest)
		return
	}

	var params NewMatchParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		utils.HTTPJsonError(w, r, "Error decoding body", err, http.StatusBadRequest)
		return
	}

	match, err := h.svc.StartNewMatch(r.Context(), userID, params.Whites)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error creating match", err, http.StatusInternalServerError)
		return
	}

	utils.HTTPJsonResponse(w, map[string]string{
		"match_id": match.ID,
	}, http.StatusCreated)
}

func (h *Handler) HandleGetMatchesByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	if userID == "" {
		utils.HTTPJsonError(w, r, "User ID not found", nil, http.StatusBadRequest)
		return
	}

	matches, err := h.svc.GetMatchesByUserID(r.Context(), userID)
	if err != nil {
		utils.HTTPJsonError(w, r, "Error fetching user match", err, http.StatusInternalServerError)
		return
	}

	utils.HTTPJsonResponse(w, MatchListToDTO(matches), http.StatusOK)
}
