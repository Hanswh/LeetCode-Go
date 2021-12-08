package leetcode

func convert(s string, numRows int) string {
	matrix := make([][]byte, numRows, numRows)
	up, down := numRows-2, 0
	for i := 0; i < len(s); {
		if down < numRows {
			matrix[down] = append(matrix[down], s[i])
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], s[i])
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	var res []byte
	for _, row := range matrix {
		for _, item := range row {
			res = append(res, item)
		}
	}
	return string(res)
}
