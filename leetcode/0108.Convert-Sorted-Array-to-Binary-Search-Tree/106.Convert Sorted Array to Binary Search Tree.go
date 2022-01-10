package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 数组有序，则可以直接找到root，递归构造即可
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	rootIndex := len(nums) / 2

	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
}
