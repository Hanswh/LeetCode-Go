package leetcode

import "sort"

// 暴力枚举
func threeSumClosest1(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sub := abs(target - res)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if abs(target-sum) < sub {
					res, sub = sum, abs(target-sum)
				}
			}
		}
	}
	return res
}

func threeSumClosest2(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sub := abs(target - res)

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
			if sub > abs(target-sum) {
				res, sub = sum, abs(target-sum)
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
