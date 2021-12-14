package leetcode

// 二分搜索，这个矩阵可以看作一个有序的一维数组，只不过需要经过坐标转换
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, low, high := len(matrix[0]), 0, len(matrix)*len(matrix[0])-1
	for low <= high {
		mid := low + (high-low)>>1
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
