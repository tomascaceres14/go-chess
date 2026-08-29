package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tomascaceres14/go-chess/api/internal/auth"
	"github.com/tomascaceres14/go-chess/api/internal/match"
	"github.com/tomascaceres14/go-chess/api/internal/middleware"
	"github.com/tomascaceres14/go-chess/api/internal/user"
)

var (
	sKey    = os.Getenv("JWT_SIGNING_KEY")
	appName = "go-chess"
)

func main() {

	sv := http.NewServeMux()

	// Auth
	tokenProvider, err := auth.NewJWTTokenProvider(sKey, appName)
	if err != nil {
		log.Fatalf("Error creating Token Provider: %v", err)
	}

	// Users
	userRepository := user.NewMemoryRepository()
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService, tokenProvider)

	// Matches
	matchRepository := match.NewMemoryRepository()
	matchService := match.NewService(matchRepository)
	matchHandler := match.NewHandler(matchService)

	// Middleware
	mw := middleware.Middleware{
		TokenProvider: tokenProvider,
		UserService:   userService,
	}

	sv.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	sv.HandleFunc("GET /users", userHandler.HandleGetUsers)
	sv.HandleFunc("POST /users", userHandler.HandleUserRegister)
	sv.HandleFunc("POST /matches", mw.Use(matchHandler.HandleNewMatch, mw.JWTAuth))
	sv.HandleFunc("GET /ws/game/{gameID}", matchHandler.HandleGameWebSocket)

	log.Printf("Server listening on port :80")
	if err := http.ListenAndServe(":80", sv); err != nil {
		log.Fatal(err)
	}
}
