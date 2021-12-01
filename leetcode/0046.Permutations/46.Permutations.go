package leetcode

func permute(nums []int) [][]int {
	chosen := make([]bool, len(nums))
	var res [][]int
	var c []int
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
		if !chosen[i] {
			chosen[i] = true
			c = append(c, num)
			dfs(nums, chosen, c, res, count+1)
			chosen[i] = false
			c = c[:len(c)-1]
		}
	}
}
