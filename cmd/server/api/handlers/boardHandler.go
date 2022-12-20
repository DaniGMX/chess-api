package handler

import (
	"fmt"
	"net/http"
)

type BoardHandler struct {
	RootURI string
}

const (
	boardRootURI = "/boards"
)

func NewBoardHandler() *BoardHandler {
	return &BoardHandler{
		RootURI: boardRootURI,
	}
}

func (bh *BoardHandler) Handle(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
fmt.Println("POST /boards")
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello word from /board")
}
