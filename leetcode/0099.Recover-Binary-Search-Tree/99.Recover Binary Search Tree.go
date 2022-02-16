package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var inorder []*TreeNode

	var inorderTravel func(node *TreeNode)
	inorderTravel = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderTravel(node.Left)
		inorder = append(inorder, node)
		inorderTravel(node.Right)
	}
	inorderTravel(root)

	for i := 0; i < len(inorder)-1; i++ {
		if inorder[i].Val > inorder[i+1].Val {
			min := inorder[i]
			for j := i + 1; j < len(inorder); j++ {
				if inorder[j].Val < min.Val {
					min = inorder[j]
				}
			}
			inorder[i].Val, min.Val = min.Val, inorder[i].Val
		}
	}
}
