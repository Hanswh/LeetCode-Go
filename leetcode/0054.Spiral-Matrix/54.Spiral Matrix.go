package leetcode

// 思路是模拟
// 通过控制i,j按照题意顺序依次读出
func spiralOrder(matrix [][]int) []int {
	// 边界，表示不能读取的部分
	// 初始条件如下
	// 上边界		-1
	// 下边界		m
	// 左边界		-1
	// 右边界		n
	leftBorder, rightBorder, upBorder, downBorder := -1, len(matrix[0]), -1, len(matrix)

	// 探针(i,j)
	// 	螺旋遍历从左上角开始，依次按向右、向下、向左、向上遍历，转弯的条件是探针走到了边界上
	// 	探索完毕 == 左右边界相遇或上下边界相遇
	// 	初始情况我们可以假设探针刚刚经历完依次向上遍历，停在(-1,-1)的位置
	// 方向d
	// 	0，1，2，3分别表示右，下，左，上
	i, j, d := -1, -1, 0
	var res []int
	for leftBorder != rightBorder-1 && upBorder != downBorder-1 {
		switch d {
		case 0:
			// 向右
			// 上次向上探索走到了边界，重置探针位置
			i++
			j++
			for j < rightBorder {
				res = append(res, matrix[i][j])
				j++
			}
			// 向右遍历，上边界下移
			upBorder++
		case 1:
			i++
			j--
			for i < downBorder {
				res = append(res, matrix[i][j])
				i++
			}
			rightBorder--
		case 2:
			i--
			j--
			for j > leftBorder {
				res = append(res, matrix[i][j])
				j--
			}
			downBorder--
		case 3:
			i--
			j++
			for i > upBorder {
				res = append(res, matrix[i][j])
				i--
			}
			leftBorder++
		}
		// 转弯
		d = (d + 1) % 4
	}
	return res
}
