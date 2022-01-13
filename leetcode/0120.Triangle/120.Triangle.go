package leetcode

import "math"

// DP，无辅助空间
// dp[row][col]表示从这个元素开始到达末端的最短路径和
// 这里直接在原数组上构造，最后一行已经满足，从倒数第二行开始向上构造
// 状态转移方程为dp[row][col] += min(triangle[row+1][col] , triangle[row+1][col+1])
func minimumTotal(triangle [][]int) int {
	if triangle == nil {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}

// DP，空间复杂度 O(n)
// dp[row][col]表示从最初到当前位置的最短路径和
// 采用滚动数组优化（按行从上到下构造，只用到前一行）
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp, minNum, index := make([]int, len(triangle[len(triangle)-1])), math.MaxInt64, 0
	for ; index < len(triangle[0]); index++ {
		dp[index] = triangle[0][index]
	}
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				// 最左边
				dp[j] += triangle[i][0]
			} else if j == len(triangle[i])-1 {
				// 最右边
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				// 中间
				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] < minNum {
			minNum = dp[i]
		}
	}
	return minNum
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
