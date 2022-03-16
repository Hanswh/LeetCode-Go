package leetcode

// dfs 超时辽
func numDistinct(s string, t string) int {
	res := 0

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if j == len(t) {
			res++
			return
		}
		for index := i + 1; index < len(s); index++ {
			if s[index] == t[j] {
				dfs(index, j+1)
			}
		}
	}
	dfs(-1, 0)
	return res
}

// dp
// dp[i][j] 表示s[:i]中可以有多少种方法组成t[:j]
// dp[i][0] = 1
// dp[0][j] = 0 (j > 0)
// 当s[i-1] == t[j-1]时	如果使用s[i-1]去匹配t[j-1]，则子序列数为dp[i-1][j-1]
// 					如果不用s[i]去匹配t[j]，则子序列数为dp[i-1][j]（用i之前的数匹配j)
// 当s[i-1] != t[j-1]时	子序列数为dp[i-1][j]
// 所求为dp[len(s)][len(t)]
func numDistinct2(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(t); j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i > 0 {
				if s[i-1] == t[j-1] {
					dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(t)]
}
