package leetcode

func dominantIndex(nums []int) int {
	maxNum, secNum, index := -1, -1, 0
	for i := range nums {
		if maxNum <= nums[i] {
			maxNum, secNum, index = nums[i], maxNum, i
		} else if secNum <= nums[i] {
			secNum = nums[i]
		}
	}
	if secNum*2 <= maxNum {
		return index
	}
	return -1
}
