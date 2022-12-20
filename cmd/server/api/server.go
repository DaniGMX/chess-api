package api

import (
	"net/http"

	handlers "github.com/danigmx/chess-api/cmd/server/api/handlers"
)

type Server struct {
	listenAddress string

	boardHandler *handlers.BoardHandler
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		boardHandler:  handlers.NewBoardHandler(),
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleRoot)
	http.HandleFunc(s.boardHandler.RootURI, s.handleBoard)
	return http.ListenAndServe(s.listenAddress, nil)
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {}

func (s *Server) handleBoard(w http.ResponseWriter, r *http.Request) {
	s.boardHandler.Handle(w, r)
}
