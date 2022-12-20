package chess

import (
	"math/bits"
)

func setBit(bitboard *uint64, square int)       { *bitboard |= (uint64(1) << square) }
func getBit(bitboard uint64, square int) uint64 { return bitboard & (uint64(1) << square) }
func popBit(bitboard *uint64, square int)       { *bitboard &= bits.ReverseBytes64(uint64(1) << square) }
func isZero64(bitboard uint64) bool             { return bitboard == 0 }
func isZero(bitboard int) bool                  { return bitboard == 0 }

func countBits(bitboard uint64) int { return bits.OnesCount64(bitboard) }
func lsb(bitboard uint64) int {
	if bitboard != 0 {
		return countBits(bitboard)
	} else {
		return -1
	}
}
