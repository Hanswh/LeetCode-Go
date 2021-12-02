package leetcode

func solveNQueens(n int) [][]string {
	col, dia1, dia2 := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1)
	var boards [][]string
	dfs(col, dia1, dia2, []string{}, 0, &boards)
	return boards
}

func dfs(col, dia1, dia2 []bool, board []string, rowIndex int, boards *[][]string) {
	n := len(col)
	if rowIndex == n {
		tmp := make([]string, len(board))
		copy(tmp, board)
		*boards = append(*boards, tmp)
	}

	for c := 0; c < n; c++ {
		if !col[c] && !dia1[rowIndex+n-c-1] && !dia2[c+rowIndex] {
			col[c] = true
			dia1[rowIndex+n-c-1] = true
			dia2[c+rowIndex] = true
			b := make([]byte, n)
			for j := 0; j < n; j++ {
				if j == c {
					b[j] = 'Q'
				} else {
					b[j] = '.'
				}
			}
			board = append(board, string(b))
			dfs(col, dia1, dia2, board, rowIndex+1, boards)
			board = board[:len(board)-1]
			col[c] = false
			dia1[rowIndex+n-c-1] = false
			dia2[c+rowIndex] = false
		}
	}
}

// 解法二 二进制操作法 Signed-off-by: Hanlin Shi shihanlin9@gmail.com
func solveNQueens2(n int) (res [][]string) {
	placements := make([]string, n)
	for i := range placements {
		buf := make([]byte, n)
		for j := range placements {
			if i == j {
				buf[j] = 'Q'
			} else {
				buf[j] = '.'
			}
		}
		placements[i] = string(buf)
	}
	var construct func(prev []int)
	construct = func(prev []int) {
		if len(prev) == n {
			plan := make([]string, n)
			for i := 0; i < n; i++ {
				plan[i] = placements[prev[i]]
			}
			res = append(res, plan)
			return
		}
		occupied := 0
		for i := range prev {
			dist := len(prev) - i
			bit := 1 << prev[i]
			occupied |= bit | bit<<dist | bit>>dist
		}
		prev = append(prev, -1)
		for i := 0; i < n; i++ {
			if (occupied>>i)&1 != 0 {
				continue
			}
			prev[len(prev)-1] = i
			construct(prev)
		}
	}
	construct(make([]int, 0, n))
	return
}
