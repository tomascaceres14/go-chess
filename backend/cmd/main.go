package main

import (
	"fmt"
    "database/sql"
	_ "github.com/lib/pq"

	"github.com/tomascaceres14/go-chess/backend/internal/engine"
)

type apiConfig struct {
	//DbQueries   *database.Queries
	Platform    string
	JwtSecret   string
	PolkaAPIKey string
}

type apiError struct {
	Error string `json:"error"`
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

	db, err := sql.Open("sqlite", "../go_chess.db")
	    if err != nil {
        fmt.Println(err)
        return
    }

	defer db.Close()



	// godotenv.Load(".env")

	// dbURL := os.Getenv("DB_URL")
	// platform := os.Getenv("PLATFORM")
	// jwtSecret := os.Getenv("JWT_SECRET")

	// db, err := sql.Open("postgres", dbURL)
	// if err != nil {
	// 	log.Fatal("DB connection error")
	// }

	// dbQueries := database.New(db)

	// apiCfg := apiConfig{
	// 	DbQueries: dbQueries,
	// 	Platform:  platform,
	// 	JwtSecret: jwtSecret,
	// }

	// mux := http.NewServeMux()

}
