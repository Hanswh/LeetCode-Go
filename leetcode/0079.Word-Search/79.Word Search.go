package leetcode

// 回溯法
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	// 标记数组
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	// board[pi][pj]表示上一个搜索到的位置，index表示接下来要搜索word中的第index个字符
	var dfs func(pi, pj, index int) bool
	dfs = func(pi, pj, index int) bool {
		if index == len(word) {
			return true
		}
		// 第一个可以遍历全部
		if index == 0 {
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					if board[i][j] == word[0] {
						used[i][j] = true
						if dfs(i, j, 1) {
							return true
						}
						used[i][j] = false
					}
				}
			}
		} else {
			// 在之前的基础上向各个方向搜索
			// 向左
			if pj-1 >= 0 && !used[pi][pj-1] && board[pi][pj-1] == word[index] {
				used[pi][pj-1] = true
				if dfs(pi, pj-1, index+1) {
					return true
				}
				used[pi][pj-1] = false
			}
			// 向右
			if pj+1 < n && !used[pi][pj+1] && board[pi][pj+1] == word[index] {
				used[pi][pj+1] = true
				if dfs(pi, pj+1, index+1) {
					return true
				}
				used[pi][pj+1] = false
			}
			// 向上
			if pi-1 >= 0 && !used[pi-1][pj] && board[pi-1][pj] == word[index] {
				used[pi-1][pj] = true
				if dfs(pi-1, pj, index+1) {
					return true
				}
				used[pi-1][pj] = false
			}
			// 向下
			if pi+1 < m && !used[pi+1][pj] && board[pi+1][pj] == word[index] {
				used[pi+1][pj] = true
				if dfs(pi+1, pj, index+1) {
					return true
				}
				used[pi+1][pj] = false
			}
		}
		return false
	}
	return dfs(-1, -1, 0)
}
