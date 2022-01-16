package leetcode

// 贪心
// 只要后一天价格高，就买入卖出
func maxProfit(prices []int) int {
	profit := 0
	for buy := 0; buy < len(prices)-1; buy++ {
		if prices[buy+1] > prices[buy] {
			profit += prices[buy+1] - prices[buy]
		}
	}
	return profit
}

// dp
// 用dp[i][0]表示第i天交易完后不持有股票的最大收益
// 用dp[i][1]表示第i天交易完后持有股票的最大收益
// dp[0][0] = 0		dp[0][1] = -prices[0]
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])，第i天不持有股票的最大收益为前一天不持有股票的最大收益、或持有然后今天卖出后的收益中的较大者
// dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])，第i天持有股票的最大收益为前一天持有股票的最大收益、或不持有然后今天买入后的收益中的较大者
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// dp
// 滚动数组优化
func maxProfit3(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}
	return dp0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
