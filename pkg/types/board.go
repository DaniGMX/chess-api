package types

type BoardState struct {
	Squares         [8][8]string      `json:squares`
	SideToPlay      string            `json:sideToPlay`
	CastlingRights  map[string]string `json:castlingRights`
	OpenEnPassant   int               `json:openEnPassant`
	HalfMoveClock   int               `json:halfMoveClock`
	FullMoveCounter int               `json:fullMoveCounter`
}
