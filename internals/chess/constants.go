package chess

const (
	A8 square = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	NoSquare
)

const (
	WKC castling = 1
	WQC castling = 2
	BKC castling = 4
	BQC castling = 8
)

const (
	White color = iota
	Black
	Both
)

const (
	Rook piece = iota
	Bishop
)

const (
	WP piece = iota
	WN
	WB
	WR
	WQ
	WK
	BP
	BN
	BB
	BR
	BQ
	BK
	NoPiece
)

const ASCII_PIECES = "PNBRQKpnbrqk"
const UNICODE_PICES = "♙♘♗♖♕♔♗♘♙"

var PieceFromChar map[byte]int = map[byte]int{
	'P': WP,
	'N': WN,
	'B': WB,
	'R': WR,
	'Q': WQ,
	'K': WK,
	'p': BP,
	'n': BN,
	'b': BB,
	'r': BR,
	'q': BQ,
	'k': BK,
}

const EmptyFEN = "8/8/8/8/8/8/8/8 b - -"
const StartFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
