package leetcode

func maxSubArray(nums []int) int {
	maxSub := nums[0]
	tmpSub := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		if tmpSub+nums[i] > nums[i] {
			tmpSub = tmpSub + nums[i]
		} else {
			tmpSub = nums[i]
		}
		if tmpSub > maxSub {
			maxSub = tmpSub
		}
	}
	return maxSub
}
