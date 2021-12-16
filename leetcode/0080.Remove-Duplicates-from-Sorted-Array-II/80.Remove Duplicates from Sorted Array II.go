package leetcode

// 双指针
// 快指针从头遍历数组，慢指针指向修改后的数组的末端，当慢指针指向倒数第二个数与快指针指向的数不相等时，才移动慢指针，同时赋值慢指针
func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
