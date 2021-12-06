package leetcode

// dp
// dp[i][j]表示到达(i,j)有多少种不同路径
// 初始情况下，若obstacleGrid[0][0] != 1，则dp[0][0] = 1。
// 又因为只能向下向右走，有dp[i][j] = dp[i][j-1] + dp[i-1][j]，第一行和第一列特殊处理
// 遍历依次更新dp数组，dp[m-1][n-1]即为所求
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if dp[0][i-1] != 0 && obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		if dp[i-1][0] != 0 && obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// dfs 超时
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	count := 0
	inPath := make([][]bool, len(obstacleGrid))
	for i := range inPath {
		inPath[i] = make([]bool, len(obstacleGrid[0]))
	}
	if obstacleGrid[0][0] == 0 {
		inPath[0][0] = true
		dfs(obstacleGrid, inPath, 0, 0, &count)
	}
	return count
}

func dfs(obstacleGrid [][]int, inPath [][]bool, i int, j int, count *int) {
	if i == len(obstacleGrid)-1 && j == len(obstacleGrid[0])-1 {
		*count++
		return
	}
	// right
	if j+1 < len(obstacleGrid[0]) && obstacleGrid[i][j+1] == 0 && !inPath[i][j+1] {
		inPath[i][j+1] = true
		dfs(obstacleGrid, inPath, i, j+1, count)
		inPath[i][j+1] = false
	}
	if i+1 < len(obstacleGrid) && obstacleGrid[i+1][j] == 0 && !inPath[i+1][j] {
		inPath[i+1][j] = true
		dfs(obstacleGrid, inPath, i+1, j, count)
		inPath[i+1][j] = false
	}
}
