package mediumLeetCode

import (
	"math"
	"strconv"
)

func MinOperations(boxes string) []int {
	boxesLen := len(boxes)
	retArr := make([]int, boxesLen)
	for boxIndex := 0; boxIndex < boxesLen; boxIndex++ {
		for otherBoxIndex := 0; otherBoxIndex < boxesLen; otherBoxIndex++ {
			if boxIndex != otherBoxIndex {
				num, _ := strconv.Atoi(string(boxes[otherBoxIndex]))
				retArr[boxIndex] += num*int(math.Abs(float64(otherBoxIndex - boxIndex)))
			}
		}
	}
	return retArr
}
