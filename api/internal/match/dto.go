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
	UserID  string `json:"user_id,omitempty"`
	Command string `json:"cmd"`
	Valid   bool   `json:"valid"`
	Status  string `json:"status"`
	Fen     string `json:"fen"`
	Error   string `json:"error,omitempty"`
}

type MatchDTO struct {
	ID          string   `json:"id"`
	OwnerID     string   `json:"owner_id"`
	OpponentID  string   `json:"opponent_id"`
	Status      string   `json:"status"`
	OwnerWhite  bool     `json:"owner_white"`
	FEN         string   `json:"fen"`
	MoveHistory []string `json:"move_history"`
}

func MatchToDTO(m *Match) *MatchDTO {
	return &MatchDTO{
		ID:          m.ID,
		OwnerID:     m.WhitesID,
		OpponentID:  m.BlacksID,
		Status:      m.Status,
		OwnerWhite:  m.OwnerWhite,
		FEN:         m.FEN,
		MoveHistory: m.MoveHistory,
	}
}

func MatchListToDTO(list []*Match) []*MatchDTO {
	result := make([]*MatchDTO, len(list))
	for i, v := range list {
		result[i] = MatchToDTO(v)
	}
	return result
}
