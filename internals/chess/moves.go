package chess

func encodeMove(source, target square, piece, promotion piece, captureFlag, doublePawnPushFlag, enpassandFlag, castlesFlag flag) move {
	return source | (target << 6) | (piece << 12) | (promotion << 16) | (intFlag(captureFlag) << 20) | (intFlag(doublePawnPushFlag) << 21) | (intFlag(enpassandFlag) << 22) | (intFlag(castlesFlag) << 23)
}

func decodeSourceSquare(m move) square     { return m & 0x3F }
func decodeTargetSquare(m move) square     { return (m & 0xFC0) >> 6 }
func decodePieceMoving(m move) piece       { return (m & 0xF000) >> 12 }
func decodePromotingPiece(m move) piece    { return (m & 0xF0000) >> 16 }
func decodeCaptureFlag(m move) flag        { return ((m & 0x100000) >> 20) != 0 }
func decodeDoublePawnPushFlag(m move) flag { return ((m & 0x200000) >> 21) != 0 }
func decodeEnpassantFlag(m move) flag      { return ((m & 0x400000) >> 22) != 0 }
func decodeCaslesFlag(m move) flag         { return ((m & 0x800000) >> 23) != 0 }

func generateMovesForPiece(b *Board, p piece) []move {
	moves := make([]move, 0)

	var source, target square
	bitboard := b.Bitboards[p]

	for !emptyBitboard(bitboard) {
		source = lsb(bitboard)
		attacks := pieceAttacks(b, p, source)

		for !emptyBitboard(attacks) {
			target = lsb(attacks)
			var occupancies uint64

			if b.Side == White {
				occupancies = b.Occupancies[White]
			} else {
				occupancies = b.Occupancies[Black]
			}

			if emptyBitboard(getBit(occupancies, target)) {
				moves = append(moves, encodeMove(source, target, p, 0, false, false, false, false))
			}

			popBit(&attacks, target)
		}

		popBit(&bitboard, source)
	}

	return moves
}

func generatePawnMoves(b *Bard, p piece, b bitboard) {
	var queen, rook, bishop, knight piece
	var opponent color

	if b.Side == White {
		queen, rook, bishop, knight = WQ, WR, WB, WN
		opponent = Black
	} else {
		queen, rook, bishop, knight = BQ, BR, BB, BN
		opponent = White
	}

	for !emptyBitboard(b) {
		source := lsb(b)
		var target square

		if b.Side == White {
			target = source - 8
		} else {
			target = source + 8
		}
	}
}
