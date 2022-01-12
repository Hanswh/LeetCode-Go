package leetcode

// 滚动数组构造
func getRow1(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		last := 0
		for j := 0; j < i+1; j++ {
			if j != 0 && j != i {
				tmp := last
				last = res[j]
				res[j] = res[j] + tmp
			} else {
				res[j] = 1
				last = 1
			}
		}
	}
	return res
}

// 递推公式
func getRow2(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
