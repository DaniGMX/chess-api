package chess

// Bitboard with all bits set except for the A file.
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
// 0 1 1 1 1 1 1 1
const notFileA Bitboard = 18374403900871474942

// Bitboard with all bits set except for the H file.
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
// 1 1 1 1 1 1 1 0
const notFileH Bitboard = 9187201950435737471

// Bitboard with all bits set except for the A and B files.
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
// 0 0 1 1 1 1 1 1
const notFilesAB Bitboard = 18229723555195321596

// Bitboard with all bits set except for the G and H files.
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
// 1 1 1 1 1 1 0 0
const notFilesGH Bitboard = 4557430888798830399

// Bishop relevant occupancy bit count for every square on board
var bishopRelevantBits [64]int = [64]int{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

// Rook relevant occupancy bit count for every square on board
var rookRelevantBits [64]int = [64]int{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

var rookMagicNumbers [64]Bitboard = [64]Bitboard{
	0xa080041440042080,
	0xa840200410004001,
	0xc800c1000200081,
	0x100081001000420,
	0x200020010080420,
	0x3001c0002010008,
	0x8480008002000100,
	0x2080088004402900,
	0x800098204000,
	0x2024401000200040,
	0x100802000801000,
	0x120800800801000,
	0x208808088000400,
	0x2802200800400,
	0x2200800100020080,
	0x801000060821100,
	0x80044006422000,
	0x100808020004000,
	0x12108a0010204200,
	0x140848010000802,
	0x481828014002800,
	0x8094004002004100,
	0x4010040010010802,
	0x20008806104,
	0x100400080208000,
	0x2040002120081000,
	0x21200680100081,
	0x20100080080080,
	0x2000a00200410,
	0x20080800400,
	0x80088400100102,
	0x80004600042881,
	0x4040008040800020,
	0x440003000200801,
	0x4200011004500,
	0x188020010100100,
	0x14800401802800,
	0x2080040080800200,
	0x124080204001001,
	0x200046502000484,
	0x480400080088020,
	0x1000422010034000,
	0x30200100110040,
	0x100021010009,
	0x2002080100110004,
	0x202008004008002,
	0x20020004010100,
	0x2048440040820001,
	0x101002200408200,
	0x40802000401080,
	0x4008142004410100,
	0x2060820c0120200,
	0x1001004080100,
	0x20c020080040080,
	0x2935610830022400,
	0x44440041009200,
	0x280001040802101,
	0x2100190040002085,
	0x80c0084100102001,
	0x4024081001000421,
	0x20030a0244872,
	0x12001008414402,
	0x2006104900a0804,
	0x1004081002402,
}

var bishopMagicNumbers [64]Bitboard = [64]Bitboard{
	0x40040822862081,
	0x40810a4108000,
	0x2008008400920040,
	0x61050104000008,
	0x8282021010016100,
	0x41008210400a0001,
	0x3004202104050c0,
	0x22010108410402,
	0x60400862888605,
	0x6311401040228,
	0x80801082000,
	0x802a082080240100,
	0x1860061210016800,
	0x401016010a810,
	0x1000060545201005,
	0x21000c2098280819,
	0x2020004242020200,
	0x4102100490040101,
	0x114012208001500,
	0x108000682004460,
	0x7809000490401000,
	0x420b001601052912,
	0x408c8206100300,
	0x2231001041180110,
	0x8010102008a02100,
	0x204201004080084,
	0x410500058008811,
	0x480a040008010820,
	0x2194082044002002,
	0x2008a20001004200,
	0x40908041041004,
	0x881002200540404,
	0x4001082002082101,
	0x8110408880880,
	0x8000404040080200,
	0x200020082180080,
	0x1184440400114100,
	0xc220008020110412,
	0x4088084040090100,
	0x8822104100121080,
	0x100111884008200a,
	0x2844040288820200,
	0x90901088003010,
	0x1000a218000400,
	0x1102010420204,
	0x8414a3483000200,
	0x6410849901420400,
	0x201080200901040,
	0x204880808050002,
	0x1001008201210000,
	0x16a6300a890040a,
	0x8049000441108600,
	0x2212002060410044,
	0x100086308020020,
	0x484241408020421,
	0x105084028429c085,
	0x4282480801080c,
	0x81c098488088240,
	0x1400000090480820,
	0x4444000030208810,
	0x1020142010820200,
	0x2234802004018200,
	0xc2040450820a00,
	0x2101021090020,
}

var castlingRights [64]int = [64]int{
	7, 15, 15, 15, 3, 15, 15, 11,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15,
	13, 15, 15, 15, 12, 15, 15, 14,
}

var pawnAttacks [2][64]Bitboard
var knightAttacks [64]Bitboard
var kingAttacks [64]Bitboard
var bishopMasks [64]Bitboard
var bishopAttacks [64][512]Bitboard
var rookMasks [64]Bitboard
var rookAttacks [64][4096]Bitboard

func initializeLeapersAttacks() {
	for square := 0; square < 64; square++ {
		pawnAttacks[White][square] = pawnAttackMaskForSideAndSquare(White, square)
		pawnAttacks[Black][square] = pawnAttackMaskForSideAndSquare(Black, square)
		knightAttacks[square] = knightAttackMaskForSquare(square)
		kingAttacks[square] = kingAttackMasForSquare(square)
	}
}

func initializeSliderAttacks() {
	for square := 0; square < 64; square++ {
		bishopMasks[square] = bishopAttackMaskForSquare(square)
		rookMasks[square] = rookAttackMaskForSquare(square)

		bishopAttackMask, rookAttackMask := bishopMasks[square], rookMasks[square]
		bishopRelevantBitsCount, rookRelevantBitsCount := countBits(bishopAttackMask), countBits(rookAttackMask)
		bishopOccupancyIndices, rookOccupancyIndices := (1 << bishopRelevantBitsCount), (1 << rookRelevantBitsCount)

		for index := 0; index < bishopOccupancyIndices; index++ {
			bishopOccupancy := setOccupancy(index, bishopRelevantBitsCount, bishopAttackMask)
			bishopMagicIndex := (bishopOccupancy * bishopMagicNumbers[square]) >> (64 - bishopRelevantBits[square])
			bishopAttacks[square][bishopMagicIndex] = bishopRealTimeAttackMask(square, bishopOccupancy)
		}

		for index := 0; index < rookOccupancyIndices; index++ {
			rookOccupancy := setOccupancy(index, rookRelevantBitsCount, rookAttackMask)
			rookMagicIndex := (rookOccupancy * rookMagicNumbers[square]) >> (64 - rookRelevantBits[square])
			rookAttacks[square][rookMagicIndex] = rookRealTimeAttacksMask(square, rookOccupancy)
		}
	}
}

func pawnAttackMaskForSideAndSquare(side Color, square Square) Bitboard {
	var panwAttacks Bitboard = 0
	var pawn Bitboard = 0

	setBit(&pawn, square)

	if side == White {
		if !isZero64((pawn >> 7) & notFileA) {
			panwAttacks |= (pawn >> 7)
		}
		if !isZero64((pawn >> 9) & notFileH) {
			panwAttacks |= (pawn >> 9)
		}
	} else {
		if !isZero64((pawn >> 7) & notFileA) {
			panwAttacks |= (pawn >> 7)
		}
		if !isZero64((pawn >> 9) & notFileH) {
			panwAttacks |= (pawn >> 9)
		}
	}

	return panwAttacks
}

func knightAttackMaskForSquare(square Square) Bitboard {
	var knightAttacks Bitboard = 0
	var knight Bitboard = 0
	setBit(&knight, square)

	if !isZero64((knight >> 17) & notFileH) {
		knightAttacks |= (knight >> 17)
	}
	if !isZero64((knight >> 15) & notFileH) {
		knightAttacks |= (knight >> 15)
	}
	if !isZero64((knight >> 10) & notFileH) {
		knightAttacks |= (knight >> 10)
	}
	if !isZero64((knight >> 6) & notFileH) {
		knightAttacks |= (knight >> 6)
	}

	if !isZero64((knight << 17) & notFileH) {
		knightAttacks |= (knight << 17)
	}
	if !isZero64((knight << 15) & notFileH) {
		knightAttacks |= (knight << 15)
	}
	if !isZero64((knight << 10) & notFileH) {
		knightAttacks |= (knight << 10)
	}
	if !isZero64((knight << 6) & notFileH) {
		knightAttacks |= (knight << 6)
	}

	return knightAttacks
}

func kingAttackMasForSquare(square Square) Bitboard {
	var kingAttacks Bitboard = 0
	var king Bitboard = 0
	setBit(&king, square)

	if !isZero64(king >> 8) {
		kingAttacks |= (king >> 8)
	}
	if !isZero64((king >> 9) & notFileH) {
		kingAttacks |= (king >> 9)
	}
	if !isZero64((king >> 7) & notFileA) {
		kingAttacks |= (king >> 7)
	}
	if !isZero64((king >> 1) & notFileH) {
		kingAttacks |= (king >> 1)
	}

	if !isZero64(king << 8) {
		kingAttacks |= (king << 8)
	}
	if !isZero64((king << 9) & notFileA) {
		kingAttacks |= (king << 9)
	}
	if !isZero64((king << 7) & notFileH) {
		kingAttacks |= (king << 7)
	}
	if !isZero64((king << 1) & notFileA) {
		kingAttacks |= (king << 1)
	}

	return kingAttacks
}

func bishopAttackMaskForSquare(square Square) Bitboard {
	var bishopAttacks Bitboard = 0
	var bishop Bitboard = 0
	setBit(&bishop, square)

	tr, tf := square/8, square%8

	for r, f := tr+1, tf+1; r <= 6 && f <= 6; r, f = r+1, f+1 {
		bishopAttacks |= 1 << (r*8 + f)
	}
	for r, f := tr-1, tf+1; r >= 1 && f <= 6; r, f = r-1, f+1 {
		bishopAttacks |= 1 << (r*8 + f)
	}
	for r, f := tr+1, tf-1; r <= 6 && f >= 1; r, f = r+1, f-1 {
		bishopAttacks |= 1 << (r*8 + f)
	}
	for r, f := tr-1, tf-1; r >= 1 && f >= 1; r, f = r-1, f-1 {
		bishopAttacks |= 1 << (r*8 + f)
	}

	return bishopAttacks
}

func rookAttackMaskForSquare(square Square) Bitboard {
	var rookAttacks Bitboard = 0
	var rook Bitboard = 0
	setBit(&rook, square)

	tr, tf := square/8, square%8

	for r := tr + 1; r <= 6; r++ {
		rookAttacks |= 1 << (r*8 + tf)
	}
	for r := tr - 1; r >= 1; r-- {
		rookAttacks |= 1 << (r*8 + tf)
	}
	for f := tf + 1; f <= 6; f++ {
		rookAttacks |= 1 << (tr*8 + f)
	}
	for f := tf - 1; f >= 1; f-- {
		rookAttacks |= 1 << (tr*8 + f)
	}
	return rookAttacks
}

func bishopAttackMaskForSquareAndOccupancy(square Square, occupancy Bitboard) Bitboard {
	occupancy &= bishopMasks[square]
	occupancy *= bishopMagicNumbers[square]
	occupancy >>= 64 - bishopRelevantBits[square]

	return bishopAttacks[square][occupancy]
}

func rookAttackMaskForSquareAndOccupancy(square Square, occupancy Bitboard) Bitboard {
	occupancy &= rookMasks[square]
	occupancy *= rookMagicNumbers[square]
	occupancy >>= 64 - rookRelevantBits[square]

	return rookAttacks[square][occupancy]
}

func queenAttackMaskForSquareAndOccupancy(square Square, occupancy Bitboard) Bitboard {
	return bishopAttackMaskForSquareAndOccupancy(square, occupancy) | rookAttackMaskForSquareAndOccupancy(square, occupancy)
}

func bishopRealTimeAttackMask(square Square, occupancy Bitboard) Bitboard {
	var bishop Bitboard = 0
	setBit(&bishop, square)
	var realTimeBishopAttacks Bitboard = 0
	tr, tf := square/8, square%8

	for r, f := tr+1, tf+1; r <= 7 && f <= 7; r, f = r+1, f+1 {
		realTimeBishopAttacks |= (1 << (r*8 + f))
		if (1<<(r*8+f))&occupancy != 0 {
			break
		}
	}
	for r, f := tr-1, tf+1; r >= 0 && f <= 7; r, f = r-1, f+1 {
		realTimeBishopAttacks |= (1 << (r*8 + f))
		if (1<<(r*8+f))&occupancy != 0 {
			break
		}
	}
	for r, f := tr+1, tf-1; r <= 7 && f >= 0; r, f = r+1, f-1 {
		realTimeBishopAttacks |= (1 << (r*8 + f))
		if (1<<(r*8+f))&occupancy != 0 {
			break
		}
	}
	for r, f := tr-1, tf-1; r >= 0 && f >= 0; r, f = r-1, f-1 {
		realTimeBishopAttacks |= (1 << (r*8 + f))
		if (1<<(r*8+f))&occupancy != 0 {
			break
		}
	}

	return realTimeBishopAttacks
}

func rookRealTimeAttacksMask(square Square, occupancy Bitboard) Bitboard {
	var rook Bitboard = 0
	setBit(&rook, square)
	var realTimeRookAttacks Bitboard = 0
	tr, tf := square/8, square%8

	for r := tr + 1; r <= 7; r++ {
		realTimeRookAttacks |= (1 << (r*8 + tf))
		if (1<<(r*8+tf))&occupancy != 0 {
			break
		}
	}
	for r := tr - 1; r >= 0; r-- {
		realTimeRookAttacks |= (1 << (r*8 + tf))
		if (1<<(r*8+tf))&occupancy != 0 {
			break
		}
	}
	for f := tf + 1; f <= 7; f++ {
		realTimeRookAttacks |= (1 << (tr*8 + f))
		if (1<<(tr*8+f))&occupancy != 0 {
			break
		}
	}
	for f := tf - 1; f >= 0; f-- {
		realTimeRookAttacks |= (1 << (tr*8 + f))
		if (1<<(tr*8+f))&occupancy != 0 {
			break
		}
	}

	return realTimeRookAttacks
}

func setOccupancy(index, maskBitCount int, attacks Bitboard) Bitboard {
	var occupancy Bitboard = 0

	for count := 0; count < maskBitCount; count++ {
		square := lsb(attacks)
		popBit(&attacks, square)

		if index&(1<<count) != 0 {
			occupancy |= (1 << square)
		}
	}

	return occupancy
}

func IsSquareAttacked(b *Board, square Square, side Color) bool {
	if side == White {
		if pawnAttacks[Black][square]&b.Bitboards[WP] != 0 {
			return true
		}
		if knightAttacks[square]&b.Bitboards[WN] != 0 {
			return true
		}
		if bishopAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[WB] != 0 {
			return true
		}
		if rookAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[WR] != 0 {
			return true
		}
		if queenAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[WQ] != 0 {
			return true
		}
		if kingAttacks[square]&b.Bitboards[WK] != 0 {
			return true
		}
	} else {
		if pawnAttacks[side][square]&b.Bitboards[BP] != 0 {
			return true
		}
		if knightAttacks[square]&b.Bitboards[BN] != 0 {
			return true
		}
		if bishopAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[BB] != 0 {
			return true
		}
		if rookAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[BR] != 0 {
			return true
		}
		if queenAttackMaskForSquareAndOccupancy(square, b.Occupancies[Both])&b.Bitboards[BQ] != 0 {
			return true
		}
		if kingAttacks[square]&b.Bitboards[BK] != 0 {
			return true
		}
	}

	return false
}

func pieceAttacks(b *Board, piece Piece, sourceSquare Square) Bitboard {
	var attacks Bitboard = 0
	side := b.Side

	switch piece {
	case WN, BN:
		attacks = knightAttacks[sourceSquare] & ^b.Occupancies[side]
	case WB, BB:
		attacks = bishopAttackMaskForSquareAndOccupancy(sourceSquare, b.Occupancies[Both]) & ^b.Occupancies[side]
	case WR, BR:
		attacks = rookAttackMaskForSquareAndOccupancy(sourceSquare, b.Occupancies[Both]) & ^b.Occupancies[side]
	case WQ, BQ:
		attacks = queenAttackMaskForSquareAndOccupancy(sourceSquare, b.Occupancies[Both]) & ^b.Occupancies[side]
	case WK, BK:
		attacks = kingAttacks[sourceSquare] & ^b.Occupancies[side]
	}
	return attacks
}
