package leetcode

/*
基于如下规则：
	1位格雷码有两个码字 0, 1
	(n+1)位格雷码中的前2^n个码字等于n位格雷码的码字，按顺序书写，加前缀0。在数值上等于n位格雷码的码字
	(n+1)位格雷码中的后2^n个码字等于n位格雷码的码字，按逆序书写，加前缀1。在数值上等于n位格雷码的码字加2^n+1
	n+1位格雷码的集合 = n位格雷码集合(顺序)加前缀0 + n位格雷码集合(逆序)加前缀1
*/
func grayCode1(n int) []int {
	ans := make([]int, 1, 1<<n)
	// 依次构造n位格雷码，对于i位格雷码，只需要逆序遍历i-1位格雷码，同时在每个码字上加2^i-1，顺序跟在i-1位格雷码后即可
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]|1<<(i-1))
		}
	}
	return ans
}

/*
从自然码构造格雷码
Todo: 补充证明
*/
func grayCode2(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i>>1 ^ i
	}
	return ans
}
