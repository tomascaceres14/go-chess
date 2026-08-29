package main

import (
	"log"
	"net/http"

	"github.com/tomascaceres14/go-chess/api/internal/match"
	"github.com/tomascaceres14/go-chess/api/internal/middleware"
	"github.com/tomascaceres14/go-chess/api/internal/user"
)

func main() {

	sv := http.NewServeMux()

	// Users
	userRepository := user.NewMemoryRepository()
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	// Matches
	matchRepository := match.NewMemoryRepository()
	matchService := match.NewService(matchRepository)
	matchHandler := match.NewHandler(matchService)

	sv.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	sv.HandleFunc("GET /users", userHandler.HandleGetUsers)
	sv.HandleFunc("POST /users", middleware.Use(userHandler.HandleUserRegister, middleware.Mid1, middleware.Mid2, middleware.Mid3))
	sv.HandleFunc("POST /matches", matchHandler.HandleNewMatch)
	sv.HandleFunc("GET /ws/game/{gameID}", matchHandler.HandleGameWebSocket)

	log.Printf("Server listening on port :80")
	if err := http.ListenAndServe(":80", sv); err != nil {
		log.Fatal(err)
	}
}
