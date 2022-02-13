package leetcode

// dp
// dp[n]表示n个不同节点可以组成不同BST的个数
// 初始情况：dp[0]和dp[1]为1
// dp[i] = dp[0] * dp[i-1] + dp[1] * dp[i-2] + ... + dp[i-1] * dp[0]，即左右子树节点数各为k, i-1-k时不同BST个数之和
func numTrees(n int) int {
	dp := make([]int, n)

	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for k := 0; k <= i-1; k++ {
			dp[i] += dp[k] * dp[i-1-k]
		}
	}

	return dp[n]
}
