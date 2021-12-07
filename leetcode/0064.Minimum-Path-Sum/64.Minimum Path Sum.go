package leetcode

// dp
// dp[i][j]表示从左上角到grid[i][j]的最小路径和
// 只能向右或向下移动构造过程如下：
// dp[0][0] = gird[0][0]，此外第一行和第一列别的dp[i][j]也可以容易求出
// 从gird[1][1]开始遍历，dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + gird[i][j]
// dp[m-1][n-1]即为所求
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// 原地dp
func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
