package leetcode

import "math"

// 蛮力
func maxProfit(prices []int) int {
	profit := 0
	for buy := 0; buy < len(prices)-1; buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			if prices[sell]-prices[buy] > profit {
				profit = prices[sell] - prices[buy]
			}
		}
	}
	return profit
}

// dp
// 保存一个元素prices[i]之前的最小值
func maxProfit2(prices []int) int {
	lowest := math.MaxInt32
	profit := math.MinInt32
	for _, price := range prices {
		lowest = min(lowest, price)
		profit = max(profit, price-lowest)
	}
	return profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
