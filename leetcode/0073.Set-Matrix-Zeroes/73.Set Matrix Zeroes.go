package leetcode

// 其实可以使用标记数组，但是这里要求原地工作
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	flagCol, flagRow := false, false
	// flagCol表示第一列是否存在0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			flagCol = true
			break
		}
	}
	// flagRow表示第一行是否存在0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			flagRow = true
			break
		}
	}
	// 用matrix[i][0]是否为0表示第i行需要置零，用matrix[0][j]是否为0表示第j行需要置零，但先不处理
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 第i行需要置零
				matrix[0][j] = 0 // 第j列需要置零
			}
		}
	}
	// 根据第一行和第一列置零元素
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 将第一行和第一列置零
	if flagRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if flagCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
