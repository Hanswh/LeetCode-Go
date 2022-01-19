package leetcode

// hash表优化暴力法
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if v, ok := m[num]; ok && i-v <= k {
			return true
		}
		m[num] = i
	}
	return false
}

// hash表滑动窗口
func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]struct{})
	for i, num := range nums {
		if i > k {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
