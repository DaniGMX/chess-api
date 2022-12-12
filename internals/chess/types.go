package chess

type Bitboard = uint64
type Square = int
type Castling = int
type Color = int
type Piece = int
type Move = int
type Flag = bool

func intFlag(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}
