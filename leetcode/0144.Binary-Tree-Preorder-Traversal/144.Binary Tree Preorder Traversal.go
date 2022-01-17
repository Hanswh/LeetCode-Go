package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，函数栈
func preorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)

	return res
}

// 手动栈
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top != nil {
			res = append(res, top.Val)
			stack = append(stack, top.Right)
			stack = append(stack, top.Left)
		}
	}

	return res
}
