package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var c []int
	var res [][]int

	sort.Ints(candidates)
	dfs(candidates, target, 0, c, &res)

	return res
}

func dfs(nums []int, target int, index int, c []int, res *[][]int) {
	for i := index; i < len(nums); i++ {
		if nums[i] == target {
			tmp := make([]int, len(c))
			copy(tmp, c)
			*res = append(*res, append(tmp, nums[i]))
		} else if nums[i] < target {
			c = append(c, nums[i])
			dfs(nums, target-nums[i], index+1, c, res)
			c = c[:len(c)-1]
		}
	}
}
