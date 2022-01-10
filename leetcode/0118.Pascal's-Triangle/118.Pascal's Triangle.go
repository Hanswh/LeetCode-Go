package leetcode

// 1
// 1 1
// 1 2 1
func generate(numRows int) [][]int {
	var res [][]int
	var row []int
	for i := 0; i < numRows; i++ {
		row = []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if len(res) > 0 { // 这里的len(res) > 0只是一次断言，可以没有
				row = append(row, res[i-1][j-1]+res[i-1][j])
			}
		}
		res = append(res, row)
	}
	return res
}
