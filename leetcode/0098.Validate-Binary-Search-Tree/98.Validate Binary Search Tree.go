package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func isValidBST(root *TreeNode) bool {
	res := true

	var preNode *TreeNode
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if preNode == nil || node.Val > preNode.Val {
			preNode = node
		} else {
			res = false
			return
		}
		inorder(node.Right)
	}
	inorder(root)

	return res
}
