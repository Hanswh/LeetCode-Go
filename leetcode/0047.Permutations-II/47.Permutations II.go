package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	chosen := make([]bool, len(nums))
	var res [][]int
	var c []int
	sort.Ints(nums)
	dfs(nums, chosen, c, &res, 0)
	return res
}

func dfs(nums []int, chosen []bool, c []int, res *[][]int, count int) {
	if count == len(nums) {
		tmp := make([]int, len(c))
		copy(tmp, c)
		*res = append(*res, tmp)
	}
	for i, num := range nums {
		// 去重：保证在当前count值相同num只被选择了一次
		// 		如果i == 0，肯定是第一次
		// 		如果nums[i] == nums[i-1] 但chosen[i-1] == false，说明nums[i-1]在本次count已经被选择过了
		//		如果nums[i] == nums[i-1] 但chosen[i-1] == true，说明nums[i-1]是在上次count被选择过的，本次还没有
		if i > 0 && !chosen[i-1] && nums[i] == nums[i-1] {
			continue
		}
		if !chosen[i] {
			chosen[i] = true
			c = append(c, num)
			dfs(nums, chosen, c, res, count+1)
			c = c[:len(c)-1]
			chosen[i] = false
		}
	}
}
