package leetcode

// 暴力
func strStr1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}

// KMP算法
// Todo: 待学习总结
func strStr2(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
