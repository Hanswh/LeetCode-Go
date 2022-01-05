package leetcode

// dp
// dp[i] 表示解码长度为n的前缀的方法数
// 初始dp[0] = 1，空字符串
// 接着对1~n依次执行两个状态转移方程：
// dp[i] += dp[i-1] 当s[i-1] != '0'，即第i个字符可以被单独解码
// dp[i] += dp[i-2] 当s[i-2:i]可以被解码
// 执行完成后，dp[n]即为所求
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
