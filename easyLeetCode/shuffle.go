package easyLeetCode

func Shuffle(nums []int, n int) []int {
	var retArr []int
	for i := 0; i < n; i++ {
		retArr = append(retArr, nums[i])
		retArr = append(retArr, nums[i+n])
	}
	return retArr
}