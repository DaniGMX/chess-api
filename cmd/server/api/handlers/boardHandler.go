package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/danigmx/chess-api/cmd/server/api/utils"
	"github.com/danigmx/chess-api/pkg/chess"
	"github.com/danigmx/chess-api/pkg/types"
)

type BoardsHandler struct {
	RootURI string
}

func NewBoardHandler() *BoardsHandler {
	return &BoardsHandler{}
}

func (bh *BoardsHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	method := request.Method

	switch method {
	case http.MethodGet:
		bh.get(writer, request)
	case http.MethodPost:
		bh.post(writer, request)
	default:
		http.NotFound(writer, request)
	}
}

func (bh *BoardsHandler) get(responseWriter http.ResponseWriter, request *http.Request) {
	fen := strings.TrimPrefix(request.URL.Path, "/boards/")
	fen = strings.ReplaceAll(fen, "_", " ")
	if utils.CheckRegexFen(fen) {
		board := types.BoardStateFromChessBoard(chess.ParseFEN(fen))
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(responseWriter).Encode(board)
	}
}

func (bh *BoardsHandler) post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /boards")
}
