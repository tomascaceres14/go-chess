package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tomascaceres14/go-chess/api/internal/auth"
	"github.com/tomascaceres14/go-chess/api/internal/database/generated"
	"github.com/tomascaceres14/go-chess/api/internal/match"
	"github.com/tomascaceres14/go-chess/api/internal/middleware"
	"github.com/tomascaceres14/go-chess/api/internal/token"
	"github.com/tomascaceres14/go-chess/api/internal/user"
)

var (
	sKey        = os.Getenv("JWT_SIGNING_KEY")
	portHTTP    = os.Getenv("PORT_HTTP")
	appName     = "go-chess"
	databaseUrl = os.Getenv("DATABASE_URL")
)

func main() {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	defer conn.Close(ctx)

	queries := generated.New(conn)

	sv := http.NewServeMux()

	// Tokens
	tokenProvider, err := token.NewJWTTokenProvider(sKey, appName)
	if err != nil {
		log.Fatalf("Error creating Token Provider: %v", err)
	}

	// Matches
	//matchRepository := match.NewMemoryRepository()
	matchRepository := match.NewPostgresRepository(queries)
	matchService := match.NewService(matchRepository)
	matchHandler := match.NewHandler(matchService)

	// Users
	//userRepository := user.NewMemoryRepository()
	userRepository := user.NewPostgresRepository(queries)
	userService := user.NewService(userRepository, matchService)
	userHandler := user.NewHandler(userService, tokenProvider)

	// Auth
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService, tokenProvider)

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
	sv.HandleFunc("GET /users/matches", mw.Use(matchHandler.HandleGetMatchesByUser, mw.JWTAuth))

	// Matches
	sv.HandleFunc("POST /matches", mw.Use(matchHandler.HandleNewMatch, mw.JWTAuth))
	sv.HandleFunc("GET /matches", mw.Use(matchHandler.HandleGetMatchesByUser, mw.JWTAuth))

	// WS
	sv.HandleFunc("GET /ws/match/{matchID}", mw.Use(matchHandler.HandleGameWebSocket, mw.JWTAuth))

	log.Printf("Server listening on port %s", portHTTP)
	if err := http.ListenAndServe(portHTTP, sv); err != nil {
		log.Fatal(err)
	}
}
