package leetcode

import "strconv"

// 暴力搜索
func isValidSudoku1(board [][]byte) bool {
	// row
	for i := 0; i < 9; i++ {
		tmp := [10]bool{}
		for j := 0; j < 9; j++ {
			cellVal := string(board[i][j])
			if cellVal != "." {
				index, _ := strconv.Atoi(cellVal)
				if tmp[index] {
					return false
				}
				tmp[index] = true
			}
		}
	}

	// column
	for i := 0; i < 9; i++ {
		tmp := [10]bool{}
		for j := 0; j < 9; j++ {
			cellVal := string(board[j][i])
			if cellVal != "." {
				index, _ := strconv.Atoi(cellVal)
				if tmp[index] {
					return false
				}
				tmp[index] = true
			}
		}
	}

	// box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]bool{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := string(board[ii][jj])
					if cellVal != "." {
						index, _ := strconv.Atoi(cellVal)
						if tmp[index] {
							return false
						}
						tmp[index] = true
					}
				}
			}
		}
	}
	return true
}

// 添加缓存
func isValidSudoku(board [][]byte) bool {
	rowBuf, columnBuf, boxBuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowBuf[i] = make([]bool, 9)
		columnBuf[i] = make([]bool, 9)
		boxBuf[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				index := board[i][j] - '0' - byte(1)
				// check row
				if rowBuf[i][index] || columnBuf[j][index] || boxBuf[i/3*3+j/3][index] {
					return false
				}
				rowBuf[i][index] = true
				columnBuf[j][index] = true
				boxBuf[i/3*3+j/3][index] = true
			}
		}
	}
	return true
}
