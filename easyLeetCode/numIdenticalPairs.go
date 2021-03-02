package easyLeetCode

func NumIdenticalPairs(nums []int) int {
	retInt := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] { retInt++ }
		}
	}
	return retInt
}
