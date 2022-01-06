package leetcode

// 蛮力法
func largestRectangleArea1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		// 依次考察每个“柱”
		left, right := i, i
		// 向左扩展
		for left >= 0 && heights[left] >= heights[i] {
			left--
		}
		// 向右扩展
		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}
		maxArea = max(maxArea, (right-left-1)*heights[i])
	}
	return maxArea
}

// 单调栈优化寻找left和right的过程
func largestRectangleArea2(heights []int) int {
	maxArea := 0
	left, right := make([]int, len(heights)), make([]int, len(heights))

	var stack []int
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		} else {
			right[i] = len(heights)
		}
		stack = append(stack, i)
	}
	for i, height := range heights {
		maxArea = max(maxArea, height*(right[i]-left[i]-1))
	}

	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
