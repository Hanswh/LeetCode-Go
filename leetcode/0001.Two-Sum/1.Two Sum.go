package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, contained := m[target-num]; contained {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}
