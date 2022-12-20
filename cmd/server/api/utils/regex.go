package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckRegexFen(str string) bool {
	if strings.Contains(str, "_") {
		str = strings.ReplaceAll(str, "_", " ")
	}

	isFen, err := regexp.Match(`\s*([rnbqkpRNBQKP1-8]+\/){7}([rnbqkpRNBQKP1-8]+)\s[bw-]\s(([a-hkqA-HKQ]{1,4})|(-))\s(([a-h][36])|(-))\s\d+\s\d+\s*$`, []byte(str))
	checkErr(err)

	if isFen {
		fmt.Printf("FEN: %s\n", str)
	} else {
		fmt.Printf("ERROR: string '%s' is not a FEN string\n", str)
	}

	return isFen
}
