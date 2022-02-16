package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || isSymmetricSubTree(root.Left, root.Right)
}

func isSymmetricSubTree(left *TreeNode, right *TreeNode) bool {
	return (left == nil && right == nil) || (left != nil && right != nil && left.Val == right.Val && isSymmetricSubTree(left.Left, right.Right) && isSymmetricSubTree(left.Right, right.Left))
}
