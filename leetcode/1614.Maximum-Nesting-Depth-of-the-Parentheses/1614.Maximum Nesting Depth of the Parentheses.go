package leetcode

// 输入保证是有效的，且只求嵌套深度，则可以忽略除括号外的别的字符
// 用栈处理，但是我们只关注栈可达到的最大深度，故用一个变量维护即可
func maxDepth(s string) int {
	res, count := 0, 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
		}
		if res < count {
			res = count
		}
	}
	return res
}
