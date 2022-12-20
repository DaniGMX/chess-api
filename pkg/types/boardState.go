package types

import (
	"fmt"

	"github.com/danigmx/chess-api/pkg/chess"
)

type BoardState struct {
	Squares             [8][8]string    `json:"squares"`
	SideToPlay          string          `json:"sideToPlay"`
	CastlingRights      map[string]bool `json:"castlingRights"`
	OpenEnPassantSquare string          `json:"openEnPassantSquare"`
	HalfMoveClock       int             `json:"halfMoveClock"`
	FullMoveCounter     int             `json:"fullMoveCounter"`
}

func BoardStateFromChessBoard(board *chess.Board) *BoardState {
	return &BoardState{
		Squares:             mapSquares(board),
		SideToPlay:          mapSideToPlay(board),
		CastlingRights:      mapCastlingRights(board),
		OpenEnPassantSquare: mapOpenEnPassantSquare(board),
		HalfMoveClock:       mapHalfMoveClock(board),
		FullMoveCounter:     mapFullMoveCounter(board),
	}
}

func mapSquares(board *chess.Board) [8][8]string {
	var squares [8][8]string
	for p := chess.WP; p <= chess.BK; p++ {
		piece := string(chess.ASCII_PIECES[p])
		bitboard := board.Bitboards[p]
		for !chess.IsZero64(bitboard) {
			square := chess.LSB(bitboard)
			row, col := square/8, square%8
			squares[row][col] = piece
			chess.PopBit(&bitboard, square)
		}
	}
	return squares
}

func mapSideToPlay(board *chess.Board) string {
	if board.Side == chess.White {
		return "W"
	} else {
		return "B"
	}
}

func mapCastlingRights(board *chess.Board) map[string]bool {
	return map[string]bool{
		"K": chess.IsZero(board.OpenCastlings | chess.WKSC),
		"Q": chess.IsZero(board.OpenCastlings | chess.WQSC),
		"k": chess.IsZero(board.OpenCastlings | chess.BKSC),
		"q": chess.IsZero(board.OpenCastlings | chess.BQSC),
	}
}

func mapOpenEnPassantSquare(board *chess.Board) string {
	rank, file := "ABCDEFGH"[7-(board.OpenEnpassant/8)], board.OpenEnpassant%8+1
	return fmt.Sprintf("%s%d", string(rank), file)
}

func mapHalfMoveClock(board *chess.Board) int {
	return board.HalfMoveClock
}

func mapFullMoveCounter(board *chess.Board) int {
	return board.FullMoveNumber
}
