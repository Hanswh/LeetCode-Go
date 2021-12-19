package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(index int, subset []int)
	dfs = func(index int, subset []int) {
		if index == len(nums) {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			res = append(res, tmp)
			return
		}
		dfs(index+1, append(subset, nums[index]))
		dfs(index+1, subset)
	}
	dfs(0, []int{})
	return res
}
