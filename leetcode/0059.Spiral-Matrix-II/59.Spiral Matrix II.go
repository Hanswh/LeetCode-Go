package leetcode

// 同0054
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	leftBorder, rightBorder, upBorder, downBorder := -1, n, -1, n
	i, j, d, count := -1, -1, 0, 1
	for leftBorder != rightBorder-1 && upBorder != downBorder-1 {
		switch d {
		case 0:
			i++
			j++
			for j < rightBorder {
				res[i][j] = count
				count++
				j++
			}
			upBorder++
		case 1:
			i++
			j--
			for i < downBorder {
				res[i][j] = count
				count++
				i++
			}
			rightBorder--
		case 2:
			j--
			i--
			for j > leftBorder {
				res[i][j] = count
				count++
				j--
			}
			downBorder--
		case 3:
			i--
			j++
			for i > upBorder {
				res[i][j] = count
				count++
				i--
			}
			leftBorder++
		}
		d = (d + 1) % 4
	}
	return res
}
