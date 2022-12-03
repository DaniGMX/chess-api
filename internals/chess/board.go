package chess

import (
	"log"
	"strconv"
	"strings"
)

type Board struct {
	Bitboards      []uint64
	Occupancies    []uint64
	Side           int
	OpenEnpassant  int
	OpenCastlings  int
	HalfMoveClock  int
	FullMoveNumber int
}

func ParseFEN(fen string) *Board {
	var board Board
	p := 0

	board.Bitboards = make([]uint64, 12)
	board.Occupancies = make([]uint64, 3)
	board.Side = White
	board.OpenEnpassant = NoSquare
	board.OpenCastlings = 0

	parsePiecePlacement(fen, &board, &p)
	parseSideToPlay(fen, &board, &p)
	parseOpenCastlings(fen, &board, &p)
	parseOpenEnpassant(fen, &board, &p)
	parseMoveCounters(fen, &board, &p)

	return &board
}

func parseMoveCounters(fen string, board *Board, p *int) {
	tail := fen[*p:]
	if len(tail) == 0 {
		board.HalfMoveClock, board.FullMoveNumber = 0, 0
	}
	moves := strings.Split(tail, " ")
	halfMoveClock, err := strconv.Atoi(moves[0])
	if err != nil {
		log.Fatal(err)
	}
	fullMoveCounter, err := strconv.Atoi(moves[1])
	if err != nil {
		log.Fatal(err)
	}
	board.HalfMoveClock, board.FullMoveNumber = halfMoveClock, fullMoveCounter
}

func parseOpenEnpassant(fen string, board *Board, p *int) {
	if fen[*p] != '-' {
		file := (int)(fen[*p] - 'a')
		rank := 8 - (int)(fen[*p+1]-'0')
		board.OpenEnpassant = rank*8 + file
	} else {
		board.OpenEnpassant = NoSquare
	}
}

func parseOpenCastlings(fen string, board *Board, p *int) {
	for fen[*p] != ' ' {
		switch fen[*p] {
		case 'K':
			board.OpenCastlings |= WKC
		case 'Q':
			board.OpenCastlings |= WQC
		case 'k':
			board.OpenCastlings |= BKC
		case 'q':
			board.OpenCastlings |= BQC
		case '-':
		}
		*p++
	}

	*p++
}

func parseSideToPlay(fen string, board *Board, p *int) {
	if fen[*p] == 'w' {
		board.Side = White
	}
	if fen[*p] == 'b' {
		board.Side = Black
	}

	*p += 2
}

func parsePiecePlacement(fen string, board *Board, p *int) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; rank++ {
			square := rank*8 + file
			if (fen[*p] >= 'a' && fen[*p] <= 'z') || (fen[*p] >= 'A' && fen[*p] <= 'Z') {
				piece := PieceFromChar[fen[*p]]
				SetBit(&board.Bitboards[piece], square)
				*p++
			}
			if fen[*p] >= '0' && fen[*p] <= '9' {
				offset := (int)(fen[*p] - '0')
				piece := NoPiece
				for b := WP; b < BK; b++ {
					if !EmptyBitboard(GetBit(board.Bitboards[b], square)) {
						piece = b
					}
				}
				if piece == NoPiece {
					file--
				}
				file += offset
				*p++
			}
			if fen[*p] == '/' {
				*p++
			}
		}
	}

	for wpiece, bpiece := WP, BP; wpiece <= WK && bpiece <= BK; wpiece, bpiece = wpiece+1, bpiece+1 {
		board.Occupancies[White] = board.Occupancies[White] | board.Bitboards[wpiece]
		board.Occupancies[Black] = board.Occupancies[Black] | board.Bitboards[bpiece]
	}

	board.Occupancies[Both] = board.Occupancies[White] | board.Occupancies[Black]
	*p++
}
