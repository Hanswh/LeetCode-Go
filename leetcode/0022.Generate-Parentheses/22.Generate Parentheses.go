package leetcode

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(leftCount int, rightCount int, subRes string)
	dfs = func(leftCount int, rightCount int, subRes string) {
		// 当右括号数量为n时，搜索到的括号字符串subRes已经符合条件
		if rightCount == n {
			res = append(res, subRes)
			return
		}
		// 不符合条件则继续在subRes上添加括号
		// 如果左括号的数量小于n，则可以添加左括号
		if leftCount < n {
			dfs(leftCount+1, rightCount, subRes+"(")
		}
		// 如果右括号的数量小于左括号的，则可以添加右括号
		if rightCount < leftCount {
			dfs(leftCount, rightCount+1, subRes+")")
		}
	}
	dfs(0, 0, "")

	return res
}
