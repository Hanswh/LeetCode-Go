package leetcode

import "sort"

// 贪心
// 要使和最大，可以每次都取反最小的
func largestSumAfterKNegations1(nums []int, k int) int {
	for k > 0 {
		min := 0
		for i := range nums {
			if nums[min] > nums[i] {
				min = i
			}
		}
		nums[min] = -nums[min]
		k--
	}

	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

// 排序+贪心
// 排序后，最小无需每次遍历全部
func largestSumAfterKNegations2(nums []int, k int) int {
	sort.Ints(nums)
	var i int

	// 取反所有负数
	for i = 0; i < len(nums); i++ {
		if k > 0 {
			if nums[i] < 0 {
				nums[i] = -nums[i]
				k--
			} else {
				break
			}
		} else {
			break
		}
	}
	// 找到当前的最小非负数
	for i == len(nums) || (i > 0 && nums[i] > nums[i-1]) {
		i--
	}
	// 反复取反直至k==0
	for k > 0 {
		nums[i] = -nums[i]
		k--
	}

	// 求和
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
