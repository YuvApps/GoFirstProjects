package dailyChallengeLeetCode

import (
	"math"
)

func DistributeCandies(candyType []int) int {
	set := map[int]bool{}
	for i := 0; i < len(candyType); i++ {
		_, ok := set[candyType[i]]
		if ok == false {
			set[candyType[i]] = true
		}
	}
	return int(math.Min(float64(len(candyType)/2), float64(len(set))))
}
