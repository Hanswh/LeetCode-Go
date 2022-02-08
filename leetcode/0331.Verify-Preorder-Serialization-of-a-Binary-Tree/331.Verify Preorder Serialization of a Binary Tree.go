package leetcode

// 9,3,4,#,#,1,#,#,2,#,6,#,#
func isValidSerialization(preorder string) bool {
	stack := []int{1}
	var idx int
	for idx = 0; idx < len(preorder); idx++ {
		if preorder[idx] == '#' {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2)
		}
		if len(stack) == 0 {
			break
		}
	}
	return len(stack) == 0 && idx == len(preorder)
}
