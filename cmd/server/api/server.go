package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	listenAddress string
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleRoot)
	http.HandleBoard("/board", s.handleBoard)
	return http.ListenAndServe(s.listenAddress, nil)
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	}
	fmt.Println("Hello world!")
}

func (s *Server) handleBoard(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case http.MethodGet:

	}
}
