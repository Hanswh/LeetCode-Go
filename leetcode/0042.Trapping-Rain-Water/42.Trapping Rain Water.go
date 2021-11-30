package leetcode

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// {0,1,0,2,1,0,1,3,2,1,2,1}
// 对于位置i能不能上的水量，可以根据左右”围墙“的高度判断
// 形成围墙的条件：
// 		i右侧存在height[r] > height[i] && i左侧存在height[l] > height[i]
// 找到所有围墙，分别取最大值，则i上水量为：min(height[r], height[l]) - height[i]，否则水量为0
// 暴力求解过程如下：
// 		对于每个i，分别向左向右寻找最高的围墙（左侧最大值和右侧最大值）
// 		计算水量，增加到总量上
func trap1(height []int) int {
	res := 0
	for i := range height {
		leftMax := height[i]
		rightMax := height[i]
		// 寻找左侧最大值
		for l := i - 1; l >= 0; l-- {
			if height[l] > leftMax {
				leftMax = height[l]
			}
		}
		// 寻找右侧最大值
		for r := i + 1; r < len(height); r++ {
			if height[r] > rightMax {
				rightMax = height[r]
			}
		}
		// 计算水量
		res += min(leftMax, rightMax) - height[i]
	}
	return res
}

// 优化
// 不妨先遍历把每个位置的左最大和右最大求出来，避免之后重复遍历
func trap(height []int) int {
	res := 0

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	for i := 1; i < len(height); i++ {
		if height[i-1] > leftMax[i-1] {
			leftMax[i] = height[i-1]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if height[i+1] > rightMax[i+1] {
			rightMax[i] = height[i+1]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	for i := range height {
		if leftMax[i] > height[i] && rightMax[i] > height[i] {
			res += min(leftMax[i], rightMax[i]) - height[i]
		}
	}

	return res
}

func trap3(height []int) int {
	res := 0
	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right++
		}
	}
	return res
}
