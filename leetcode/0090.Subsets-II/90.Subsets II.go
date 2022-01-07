package leetcode

import "sort"

// dfs
// 问题在于如何去重
func subsetsWithDup(nums []int) [][]int {
	// 先排序，保证相同元素全部相邻
	sort.Ints(nums)
	var subsets [][]int
	var subset []int

	// 对每个元素选择/不选择来生成子集
	var dfs func(index int, lastSelected bool) // flag表示上一个元素是否被选择
	dfs = func(index int, lastSelected bool) {
		if index == len(nums) {
			subsets = append(subsets, append([]int{}, subset...))
			return
		}
		dfs(index+1, false)
		// 去重：只有此次元素和前一个不同，或相同但前一个也被选择了时，才选择此次元素
		if index > 0 && nums[index-1] == nums[index] && !lastSelected {
			return
		}
		subset = append(subset, nums[index])
		dfs(index+1, true)
		subset = subset[:len(subset)-1]
	}
	dfs(0, false)
	return subsets
}
