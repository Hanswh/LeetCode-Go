package leetcode

func nextPermutation(nums []int) {
	// 如何找到下一个更大的排列
	// 原地修改
	/*
		[1,2,3,4,5,7,6,3,2,2,1]
		[1,2,3,4,6,1,2,2,3,5,7]

		[1,2,3]	交换2和3
		[1,3,2] 交换1和2，把1原来位置后面的变成升序
		[2,1,3] 交换1和3
		[2,3,1] 交换2和3，把2原来位置后面的变成升序
		[3,1,2] 交换1和2
		[3,2,1]
	*/
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j = i+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
