package easyLeetCode

import (
	"strings"
)

func NumJewelsInStones(jewels string, stones string) int {
	retInt := 0
	for charIndex := 0; charIndex < len(jewels); charIndex++ {
		retInt += strings.Count(stones, string(jewels[charIndex]))
	}
	return retInt
}
