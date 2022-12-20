package chess

func encodeMove(source, target Square, piece, promotion Piece, captureFlag, doublePawnPushFlag, enpassandFlag, castlesFlag Flag) Move {
	return source | (target << 6) | (piece << 12) | (promotion << 16) | (intFlag(captureFlag) << 20) | (intFlag(doublePawnPushFlag) << 21) | (intFlag(enpassandFlag) << 22) | (intFlag(castlesFlag) << 23)
}

func decodeSourceSquare(m Move) Square     { return m & 0x3F }
func decodeTargetSquare(m Move) Square     { return (m & 0xFC0) >> 6 }
func decodePieceMoving(m Move) Piece       { return (m & 0xF000) >> 12 }
func decodePromotingPiece(m Move) Piece    { return (m & 0xF0000) >> 16 }
func decodeCaptureFlag(m Move) Flag        { return ((m & 0x100000) >> 20) != 0 }
func decodeDoublePawnPushFlag(m Move) Flag { return ((m & 0x200000) >> 21) != 0 }
func decodeEnpassantFlag(m Move) Flag      { return ((m & 0x400000) >> 22) != 0 }
func decodeCaslesFlag(m Move) Flag         { return ((m & 0x800000) >> 23) != 0 }

func generateMovesForPiece(b *Board, p Piece) []Move {
	moves := make([]Move, 0)

	var source, target Square
	bitboard := b.Bitboards[p]

	for !IsZero64(bitboard) {
		source = LSB(bitboard)
		attacks := pieceAttacks(b, p, source)

		for !IsZero64(attacks) {
			target = LSB(attacks)
			var occupancies uint64

			if b.Side == White {
				occupancies = b.Occupancies[White]
			} else {
				occupancies = b.Occupancies[Black]
			}

			if IsZero64(GetBit(occupancies, target)) {
				moves = append(moves, encodeMove(source, target, p, 0, false, false, false, false))
			}

			PopBit(&attacks, target)
		}

		PopBit(&bitboard, source)
	}

	return moves
}

func generatePawnMoves(board *Board, pawn Piece, bitboard Bitboard) []Move {
	var queen, rook, bishop, knight Piece
	var opponent Color

	moves := make([]Move, 0)

	if board.Side == White {
		queen, rook, bishop, knight = WQ, WR, WB, WN
		opponent = Black
	} else {
		queen, rook, bishop, knight = BQ, BR, BB, BN
		opponent = White
	}

	for !IsZero64(bitboard) {
		source := LSB(bitboard)
		var target Square

		if board.Side == White {
			target = source - 8
		} else {
			target = source + 8
		}

		var canQuietMove bool

		if board.Side == White {
			canQuietMove = !(target < A8) && IsZero64(GetBit(board.Occupancies[Both], target))
		} else {
			canQuietMove = !(target < A8) && IsZero64(GetBit(board.Occupancies[Both], target))
		}

		if canQuietMove {
			var isPromotion bool

			if board.Side == White {
				isPromotion = source >= A7 && source <= H7
			} else {
				isPromotion = source >= A2 && source <= H2
			}

			if isPromotion {
				moves = append(moves, encodeMove(source, target, pawn, queen, false, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, rook, false, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, bishop, false, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, knight, false, false, false, false))
			} else {
				moves = append(moves, encodeMove(source, target, pawn, 0, false, false, false, false))

				var canDoublePawnPush bool

				if board.Side == White {
					target -= 8
					canDoublePawnPush = source >= A2 && source <= H2 && IsZero64(GetBit(board.Occupancies[Both], target))
				} else {
					canDoublePawnPush = source >= A7 && source <= H7 && IsZero64(GetBit(board.Occupancies[Both], target))

				}

				if canDoublePawnPush {
					moves = append(moves, encodeMove(source, target, pawn, 0, false, true, false, false))
				}
			}
		}

		attacks := pawnAttacks[board.Side][source] & board.Occupancies[opponent]

		for !IsZero64(attacks) {
			target := LSB(attacks)

			var isPromotion bool

			if board.Side == White {
				isPromotion = source >= A7 && source <= H7
			} else {
				isPromotion = source >= A2 && source <= H2
			}

			if isPromotion {
				moves = append(moves, encodeMove(source, target, pawn, queen, true, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, rook, true, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, bishop, true, false, false, false))
				moves = append(moves, encodeMove(source, target, pawn, knight, true, false, false, false))
			} else {
				moves = append(moves, encodeMove(source, target, pawn, 0, true, false, false, false))
			}

			PopBit(&attacks, target)
		}

		if board.OpenEnpassant != NoSquare {
			openEnpassantAttackMask := pawnAttacks[board.Side][source] & (1 << board.OpenEnpassant)

			if !IsZero64(openEnpassantAttackMask) {
				target := LSB(openEnpassantAttackMask)
				moves = append(moves, encodeMove(source, target, pawn, 0, true, false, true, false))
			}
		}

		PopBit(&bitboard, source)
	}

	return moves
}

func generateCastlingMoves(board *Board, king Piece) []Move {
	moves := make([]Move, 0)

	var openKingSideCastle, openQueenSideCastle, kingSideConnected, queenSideConnected, safeKingSide, safeQueenSide bool
	var source, kingSideTarget, queenSideTarget Square

	if board.Side == White {
		openKingSideCastle = !IsZero(board.OpenCastlings & WKSC)
		openQueenSideCastle = !IsZero(board.OpenCastlings & WQSC)
		kingSideConnected = !IsZero64(GetBit(board.Occupancies[Both], F1)) && !IsZero64(GetBit(board.Occupancies[Both], G1))
		queenSideConnected = !IsZero64(GetBit(board.Occupancies[Both], B1)) && !IsZero64(GetBit(board.Occupancies[Both], C1))
		safeKingSide = IsSquareAttacked(board, E1, Black) && !IsSquareAttacked(board, F1, Black)
		safeKingSide = IsSquareAttacked(board, E1, Black) && !IsSquareAttacked(board, D1, Black)
		source = E1
		kingSideTarget = G1
		queenSideTarget = C1
	} else {
		openKingSideCastle = !IsZero(board.OpenCastlings & BKSC)
		openQueenSideCastle = !IsZero(board.OpenCastlings & BQSC)
		kingSideConnected = !IsZero64(GetBit(board.Occupancies[Both], F8)) && !IsZero64(GetBit(board.Occupancies[Both], G8))
		queenSideConnected = !IsZero64(GetBit(board.Occupancies[Both], B8)) && !IsZero64(GetBit(board.Occupancies[Both], C8))
		safeKingSide = IsSquareAttacked(board, E8, Black) && !IsSquareAttacked(board, F8, Black)
		safeKingSide = IsSquareAttacked(board, E8, Black) && !IsSquareAttacked(board, D8, Black)
		source = E8
		kingSideTarget = G8
		queenSideTarget = C8
	}

	if openKingSideCastle && !kingSideConnected && safeKingSide {
		moves = append(moves, encodeMove(source, kingSideTarget, king, 0, false, false, false, true))
	}
	if openQueenSideCastle && !queenSideConnected && safeQueenSide {
		moves = append(moves, encodeMove(source, queenSideTarget, king, 0, false, false, false, true))
	}

	return moves
}

func generatePieceMoves(board *Board, piece Piece, bitboard Bitboard) []Move {
	moves := make([]Move, 0)

	var source, target Square

	for !IsZero64(bitboard) {
		source = LSB(bitboard)
		attacks := pieceAttacks(board, piece, source)

		for !IsZero64(attacks) {
			target = LSB(attacks)

			var occupancies Bitboard

			if board.Side == White {
				occupancies = board.Occupancies[Black]
			} else {
				occupancies = board.Occupancies[White]
			}

			if IsZero64(GetBit(occupancies, target)) {
				moves = append(moves, encodeMove(source, target, piece, 0, false, false, false, false))
			} else {
				moves = append(moves, encodeMove(source, target, piece, 0, true, false, false, false))
			}

			PopBit(&attacks, target)
		}

		PopBit(&bitboard, source)
	}

	return moves
}
