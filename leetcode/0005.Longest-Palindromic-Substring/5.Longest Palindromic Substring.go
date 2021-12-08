package leetcode

// 中心扩散
// 对于每一个由相同元素组成的子串，左右同时扩散
func longestPalindrome1(s string) string {
	res := ""
	for i := range s {
		// 找到每一个由相同元素组成的子串
		if i > 0 && s[i] == s[i-1] { // 之前已经扩散过
			continue
		}
		tmp, left, right := s[i:i+1], i, i
		// 向右寻找相同元素
		for right+1 < len(s) && s[right+1] == s[left] {
			right++
			tmp = s[left : right+1]
		}
		// 扩散
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
			tmp = s[left : right+1]
		}
		// 更新最大值
		if len(res) < len(tmp) {
			res = tmp
		}
	}
	return res
}

// dp
// dp[i][j]表示从i到j是否是回文串
// dp[i][j] = (dp[i+1][j-1] || j-i < 3) && s[i] == s[j]
func longestPalindrome2(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	mI, mJ := 0, 0
	for right := 0; right < len(s); right++ {
		for left := right; left >= 0; left-- {
			if s[right] == s[left] && (dp[left+1][right-1] || right-left < 3) {
				dp[left][right] = true
				if right-left > mJ-mI {
					mI = left
					mJ = right
				}
			}
		}
	}
	return s[mI : mJ+1]
}

// 中心扩散法，另一种实现，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome3(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// 滑动窗口，类似解法一
func longestPalindrome4(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)>>1 + 1
		right = left
	}
	return s[pl : pr+1]
}

// Manacher's algorithm，时间复杂度 O(n)，空间复杂度 O(n)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}
	// dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
	// maxRight: 通过中心扩散的方式能够扩散的最右边的下标
	// center:   与 maxRight 对应的中心字符的下标
	// maxLen:   记录最长回文串的半径
	// begin:    记录最长回文串在起始串 s 中的起始下标
	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			// 这一行代码是 Manacher 算法的关键所在
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		// 中心扩散法更新 dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		// 记录最长回文子串的长度和相应它在原始字符串中的起点
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
		}
	}
	return s[begin : begin+maxLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
