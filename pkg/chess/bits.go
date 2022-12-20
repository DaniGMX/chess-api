package chess

import (
	"math/bits"
)

func SetBit(bitboard *uint64, square int)       { *bitboard |= (1 << square) }
func GetBit(bitboard uint64, square int) uint64 { return bitboard & (1 << square) }
func PopBit(bitboard *uint64, square int)       { *bitboard &= ^(1 << square) }
func IsZero64(bitboard uint64) bool             { return bitboard == 0 }
func IsZero(bitboard int) bool                  { return bitboard == 0 }

func CountBits(bitboard uint64) int { return bits.OnesCount64(bitboard) }
func LSB(bitboard uint64) int {
	if bitboard != 0 {
		return CountBits((bitboard & -bitboard) - 1)
	} else {
		return -1
	}
}
