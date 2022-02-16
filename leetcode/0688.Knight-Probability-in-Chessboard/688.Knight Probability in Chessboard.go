package leetcode

var dirs = [][]int{{1, 2}, {1, -2}, {2, 1}, {2, -1}, {-1, 2}, {-1, -2}, {-2, 1}, {-2, -1}}

// dp[step][i][j]
// 定义为从(i,j)开始走了step步还在棋盘上的概率
// dp[0][i][j] = 1
// 当不在棋盘上，即i < 0 || i > n-1 || j < 0 || j > n-1时，dp[step][i][j] = 0
// 对于一般的位置，dp[step][i][j] = 0.125 * Σdp[step-1][i+di][j+dj]
// 即从(i,j)出发走step步不出棋盘的概率等于所有从(i,j)可以一步到达的另一点出发走step-1步不出棋盘的概率之和
func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, k+1)
	for step := range dp {
		dp[step] = make([][]float64, n)
		for i := range dp[step] {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, dir := range dirs {
						if x, y := i+dir[0], j+dir[1]; x >= 0 && x < n && y >= 0 && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}
