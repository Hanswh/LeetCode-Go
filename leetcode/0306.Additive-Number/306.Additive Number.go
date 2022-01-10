package leetcode

import "strconv"

// dfs
// 当确认了第一个数和第二个数时，这个序列该如何切分就已经确定（每次都是n3 == n1 + n2时切)
// 要注意处理前导0的问题
func isAdditiveNumber(num string) bool {
	flag := false

	// 深度优先搜索，寻找n3并递归直至整个num可作为n3
	var dfs func(num string, n1, n2 int)
	dfs = func(num string, n1, n2 int) {
		for i := 1; i <= len(num); i++ {
			// 避免判断前导0
			if num[0] == '0' && i > 1 {
				break
			}

			n3, _ := strconv.Atoi(num[:i])
			if n3 == n1+n2 {
				// 找到了n3，如果是整个num则找到了序列，否则继续搜索
				if i == len(num) {
					flag = true
					return
				} else {
					dfs(num[i:], n2, n3)
				}
			} else if n3 > n1+n2 {
				// 剪个枝
				break
			}
		}
	}

	// 枚举所有可能的第一个和第二个数的组合
	// num[:i]为第一个数，num[i:j]为第二个数
	for i := 1; i <= len(num)-2; i++ {
		for j := i + 1; j <= len(num)-1; j++ {
			// 要保证n1和n2不能有前导0，要么长度为1，要么第一位不是0
			if (i+1 == j || num[i] != '0') && (num[0] != '0' || i == 1) {
				n1, _ := strconv.Atoi(num[:i])
				n2, _ := strconv.Atoi(num[i:j])
				dfs(num[j:], n1, n2)
			}
		}
	}
	return flag
}
