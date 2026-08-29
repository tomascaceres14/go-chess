package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tomascaceres14/go-chess/api/internal/auth"
	"github.com/tomascaceres14/go-chess/api/internal/match"
	"github.com/tomascaceres14/go-chess/api/internal/middleware"
	"github.com/tomascaceres14/go-chess/api/internal/token"
	"github.com/tomascaceres14/go-chess/api/internal/user"
)

var (
	sKey     = os.Getenv("JWT_SIGNING_KEY")
	portHTTP = os.Getenv("PORT_HTTP")
	appName  = "go-chess"
)

func main() {

	sv := http.NewServeMux()

	// Tokens
	tokenProvider, err := token.NewJWTTokenProvider(sKey, appName)
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

	// Auth
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)

	// Middleware
	mw := middleware.Middleware{
		TokenProvider: tokenProvider,
		UserService:   userService,
	}

	sv.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running!"))
	})

	// Auth endpoints
	sv.HandleFunc("POST /auth/register", authHandler.HandleRegister)
	sv.HandleFunc("POST /auth/login", authHandler.HandleLogin)

	// Users
	sv.HandleFunc("GET /users", userHandler.HandleGetUsers)

	// Matches
	sv.HandleFunc("POST /matches", mw.Use(matchHandler.HandleNewMatch, mw.JWTAuth))

	// WS
	sv.HandleFunc("GET /ws/match/{gameID}", mw.Use(matchHandler.HandleGameWebSocket, mw.JWTAuth))

	log.Printf("Server listening on port %s", portHTTP)
	if err := http.ListenAndServe(portHTTP, sv); err != nil {
		log.Fatal(err)
	}
}
