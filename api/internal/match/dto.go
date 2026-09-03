package match

const (
	MovePieceCmd = "match.cmd.move"
	MatchEndCmd  = "match.cmd.end"

	MatchBeginStatus   = "match.status.begin"
	MatchWaitingStatus = "match.status.waiting"
)

type NewMatchParams struct {
	Whites bool `json:"whites"`
}

type Move struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type GameCommand struct {
	UserID  string `json:"-"`
	Command string `json:"cmd"`
	Move    Move   `json:"move"`
}

type GameResponse struct {
	UserID  string         `json:"user_id,omitempty"`
	Command string         `json:"cmd"`
	Valid   bool           `json:"valid"`
	Status  string         `json:"status"`
	Data    map[string]any `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
	Grid    []string       `json:"grid,omitempty"`
}

type MatchDTO struct {
	ID         string `json:"id"`
	OwnerID    string `json:"owner_id"`
	OpponentID string `json:"opponent_id"`
	Status     string `json:"status"`
	OwnerWhite bool   `json:"owner_white"`
}

func MatchToDTO(m *Match) *MatchDTO {
	return &MatchDTO{
		ID:         m.ID,
		OwnerID:    m.OwnerID,
		OpponentID: m.OpponentID,
		Status:     m.Status,
		OwnerWhite: m.OwnerWhite,
	}
}

func MatchListToDTO(list []*Match) []*MatchDTO {
	result := make([]*MatchDTO, len(list))
	for i, v := range list {
		result[i] = MatchToDTO(v)
	}
	return result
}
