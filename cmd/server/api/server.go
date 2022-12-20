package api

import (
	"log"
	"net/http"

	"github.com/danigmx/chess-api/cmd/server/api/handlers"
)

type Server struct {
	port string
	mux  *http.ServeMux

	boardsHandler *handlers.BoardsHandler
	gamesHandler  *handlers.GamesHandler
}

func NewServer(port string) *Server {
	return &Server{
		port:          port,
		mux:           http.NewServeMux(),
		boardsHandler: handlers.NewBoardHandler(),
		gamesHandler:  handlers.NewGameHandler(),
	}
}

func (s *Server) Start() {
	s.mux.HandleFunc("/boards/", s.boardsHandler.Handle)

	server := &http.Server{
		Addr:    s.port,
		Handler: s.mux,
	}
	log.Fatal(server.ListenAndServe())
}
