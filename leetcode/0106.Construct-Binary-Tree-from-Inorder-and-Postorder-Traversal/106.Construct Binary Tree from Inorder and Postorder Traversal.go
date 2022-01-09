package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// left root right
	// left right root
	n := len(postorder)
	root := postorder[n-1]
	rightCount := 0
	for i := n - 1; i >= 0; i-- {
		if inorder[i] == root {
			break
		}
		rightCount++
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(inorder[:n-rightCount-1], postorder[:n-rightCount-1]),
		Right: buildTree(inorder[n-rightCount:], postorder[n-rightCount-1:n-1]),
	}
}
