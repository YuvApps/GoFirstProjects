package dailyChallengeLeetCode

func FindErrorNums(nums []int) []int {
	set := map[int]bool{}
	retArr := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		_, ok := set[nums[i]]
		if ok == false {
			set[nums[i]] = true
		} else {
			retArr[0] = nums[i]
		}
	}
	for num := 1; num <= len(nums); num++ {
		_, ok := set[num]
		if ok == false {
			retArr[1] = num
			break
		}
	}
	return retArr
}
