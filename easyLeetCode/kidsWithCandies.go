package easyLeetCode

func KidsWithCandies(candies []int, extraCandies int) []bool {
	N := len(candies)
	max := 0
	var retArr []bool
	for i := 0; i < N; i++ {
		if candies[i] > max { max = candies[i] }
	}

	for i := 0; i < N; i++ {
		if candies[i] + extraCandies >= max { retArr = append(retArr, true) } else { retArr = append(retArr, false) }
	}

	return retArr
}
