package easyLeetCode

func RunningSum(nums []int) []int {
	var retArr []int
	retArr = append(retArr, nums[0])
	for i := 1; i < len(nums); i++ {
		retArr = append(retArr, retArr[i-1] + nums[i])
	}
	return retArr
}
