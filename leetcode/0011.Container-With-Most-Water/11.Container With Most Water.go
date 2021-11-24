package leetcode

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1

	for start < end {
		width := end - start
		h := 0
		if height[start] > height[end] {
			h = height[end]
			end--
		} else {
			h = height[start]
			start++
		}

		area := width * h
		if max < area {
			max = area
		}
	}

	return max
}
