package leetcode

// dp
// 用dp[i][j]表示s1的前i个字符可以和s2的前j个字符交错组成s3
// 则dp[0][0] = true
// 状态转移方程为dp[i][j] = (dp[i-1][j] && s1[i] == s3[i+j]) || (dp[i][j-1] && s2[j] == s3[i+j])
// dp[len(s1)][len(s2)]即为所求
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if (m + n) != t {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j]
			}
			if j > 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1] || dp[i][j]
			}
		}
	}

	return dp[m][n]
}

// dp优化
// 上面的构造dp第i行是根据第i-1行构造的
func isInterleave2(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if (m + n) != t {
		return false
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= m; i++ {
		// 每次外层循环构造一行
		for j := 0; j <= n; j++ {
			// 每次内层循环会访问一遍上一行同列的值，并重新赋值
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i+j-1]
			}
			// 或访问一遍本行前一列的值，重新赋值
			if j > 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[i+j-1] || dp[j]
			}
		}
	}

	return dp[n]
}
