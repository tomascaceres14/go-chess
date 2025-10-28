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

func (*apiConfig) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	msg := "Hello World!"
	w.Write([]byte(msg))
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

	godotenv.Load("../.env")

	DB_URL := os.Getenv("DB_URL")
	DB_ENGINE := os.Getenv("DB_ENGINE")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	db, err := sql.Open(DB_ENGINE, DB_URL)
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
