package leetcode

import "sort"

// DFS
func combinationSum1(candidates []int, target int) [][]int {
	var res [][]int
	var c []int

	sort.Ints(candidates)
	dfs(candidates, target, 0, c, &res)

	return res
}

func dfs(nums []int, target int, index int, s []int, res *[][]int) {
	for i := index; i < len(nums); i++ {
		if nums[i] < target {
			s = append(s, nums[i])
			dfs(nums, target-nums[i], i, s, res)
			s = s[:len(s)-1]
		} else if nums[i] == target {
			tmp := make([]int, len(s))
			copy(tmp, s)
			*res = append(*res, append(tmp, nums[i]))
		} else {
			break
		}
	}
}
