package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	"github.com/tomascaceres14/go-chess/backend/internal/database"
	"github.com/tomascaceres14/go-chess/backend/internal/engine"
)

type apiConfig struct {
	DbQueries *database.Queries
	Platform  string
	JwtSecret string
}

type errorResponse struct {
	Error string `json:"error"`
}

func (cfg *apiConfig) HelloWorld(w http.ResponseWriter, r *http.Request) {

	if err := cfg.DbQueries.CreateGame(r.Context(), database.CreateGameParams{
		ID:          "black_v_white",
		WhitePlayer: "tomas",
		BlackPlayer: "isabela",
		Result:      sql.NullInt64{Int64: 1},
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	games, err := cfg.DbQueries.GetGames(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	print(len(games))

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(games[0].ID))
}

// Error printing for debugging
func PrintError(err error) {
	fmt.Printf("--- ERROR: %v\n", err)
}

func SwitchTurns(p1, p2 *engine.Player, white bool) *engine.Player {
	if p1.White == white {
		return p1
	}

	return p2
}

func main() {

	const port = "8080"

	godotenv.Load("../.env")

	//DB_URL := os.Getenv("DB_URL")
	//DB_ENGINE := os.Getenv("DB_ENGINE")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	db, err := sql.Open("sqlite", "./go_chess.db")
	if err != nil {
		fmt.Println("Database connection error.")
		log.Fatal(err)
	}

	defer db.Close()

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DbQueries: dbQueries,
		JwtSecret: JWT_SECRET,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/test", apiCfg.HelloWorld)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Serving on port " + port)
	log.Fatal(server.ListenAndServe())

}
