package leetcode

// dp
func jump1(nums []int) int {
	n := len(nums)
	// dp[i]表示跳到i最少需要的步数
	// 初始dp[0] = 0，其他均为一个最大值
	// 对于每个nums[i]，我们判断dp[i+1]到dp[i+nums[i]]是否大于dp[i]+1，是则更新
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && j <= i+nums[i]; j++ {
			if dp[j] > dp[i]+1 {
				dp[j] = dp[i] + 1
			}
		}
	}
	return dp[n-1]
}

// dp
func jump2(nums []int) int {
	n := len(nums)
	// dp含义和初始不变
	// 再增加一个值maxI表示在跳到i之前，可以跳跃到的最远距离，这样我们更新dp只需要将dp[maxI+1]到dp[i+nums[i]]全部更新为dp[i]+1
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = n
	}
	maxI := 0
	for i := 0; i < n; i++ {
		if maxI < i+nums[i] {
			for j := maxI + 1; j < n && j <= i+nums[i]; j++ {
				dp[j] = dp[i] + 1
			}
			maxI = i + nums[i]
		}
	}
	return dp[n-1]
}

// dp
func jump3(nums []int) int {
	// 分析上面的jump2，可以发现我们对于dp的使用只限于dp[i]，也就是当前位置的最小跳跃次数，则无需创建整个数组
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxInstance, minJumpTime := 0, 0
	needChoose := 0
	for i, num := range nums {
		if maxInstance < i+num {
			maxInstance = i + num
			if maxInstance > n-2 {
				return minJumpTime + 1
			}
		}
		if i == needChoose {
			needChoose = maxInstance
			minJumpTime++
		}
	}
	return minJumpTime
}
