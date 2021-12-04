package leetcode

// 暴力搜索，超时
func canJump1(nums []int) bool {
	return dfs(nums, 0)
}

func dfs(nums []int, p int) bool {
	if p == len(nums)-1 {
		return true
	}
	for i := 1; i <= nums[p]; i++ {
		if dfs(nums, p+i) {
			return true
		}
	}
	return false
}

// dp
// 分析发现，需要跳到“能到达最后一个位置”的位置才能成功
// 从后向前遍历，保存最靠左的“能到达最后一个位置”的位置，最后如果是0，则true
func canJump2(nums []int) bool {
	position := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= position {
			position = i
		}
	}
	return position == 0
}

// 从前向后遍历，保存能到达的最右边的位置，最后如果比n-1大，则true
func canJump3(nums []int) bool {
	n := len(nums)
	mostRight := 0
	for i := 0; i < n; i++ {
		if i <= mostRight && i+nums[i] >= mostRight {
			mostRight = i + nums[i]
		}
	}
	return mostRight >= n-1
}
