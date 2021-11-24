package leetcode

// 双指针
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	left, right := 0, 0
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}
