package leetcode

// 荷兰国旗问题
// 用双指针zero和one分别表示0和1的右边界（开放的边界）
// 对数组中元素依次遍历做如下操作：
// 	1、用2替换原nums[i]
//	2、如果原nums[i] <= 1，则使得1的右边界右移
//	3、另外如果原nums[i] == 0, 则使得0的右边界右移
// 通过以上操作，所有的元素都被2替换；同时，如果遍历过程中遇到0和1，则1的右边界后移，而如果遇到的是0，则0的右边界后移。
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n <= 1 {
			nums[one] = 1
			one++
		}
		if n == 0 {
			nums[zero] = 1
			zero++
		}
	}
}
