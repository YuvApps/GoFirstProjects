package easyLeetCode

func TwoSum(nums []int, target int) []int {

	retArr := []int{0,0}
	found := false
	for i := 0; i < len(nums) && !found; i++ {
		for j := i + 1; j < len(nums) && !found; j++ {
			if nums[i] + nums[j] == target {
				retArr[0] = i
				retArr[1] = j
				found = true
			}
		}
	}
	return retArr
}
