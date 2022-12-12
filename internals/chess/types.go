package chess

type bitboard = uint64
type square = int
type castling = int
type color = int
type piece = int
type move = int
type flag = bool

func intFlag(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}
