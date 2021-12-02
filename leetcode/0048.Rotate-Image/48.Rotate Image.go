package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	// nums[i][j]
	// 123
	// 456
	// 789
	// 00 02 22 20 = 20 00 02 22
	// 01 12 21 10 = 10 01 12 10
	// ij (j)(n-i-1) (n-i-1)(n-j-1) (n-j-1)(i)
	// 每一圈的第一行中的第一个到倒数第二个
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
