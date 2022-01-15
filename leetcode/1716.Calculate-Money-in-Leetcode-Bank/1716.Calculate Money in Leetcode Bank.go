package leetcode

func totalMoney(n int) int {
	weeks := n / 7 // 完整的周数
	days := n % 7  // 多出来的天数
	// (28 + 28 * （weeks-1) * 7) * weeks / 2
	// (1 + weeks + 1 + weeks + days - 1) * days / 2
	return (56+(weeks-1)*7)*weeks/2 + (1+weeks+weeks+days)*days/2
}
