package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归构造
// 根据当前待构造树的前序和中序序列，找出左右子树的前序和中序序列，递归调用
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 求左子树节点数
	count := 0
	for _, v := range inorder {
		if v == preorder[0] {
			break
		}
		count++
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:count+1], inorder[:count]),
		Right: buildTree(preorder[count+1:], inorder[count+1:]),
	}
}

// [3,9,20,15,6,4,7,13,8]
// [9,3,6,15,4,20,13,7,8]
// 迭代构造
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var stack []*TreeNode
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
