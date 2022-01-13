package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用栈
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	stack = append(stack, root)
	dummy := &TreeNode{}
	cur := dummy
	for len(stack) != 0 {
		// 取栈顶
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 左右不空则入栈，先序则右子树先入
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		// 加入到链表
		cur.Right = &TreeNode{
			Val: top.Val,
		}
		cur = cur.Right
	}
	root.Left = nil
	root.Right = dummy.Right.Right
}

// 使用栈，优化，不再新建节点
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	stack = append(stack, root)
	cur := &TreeNode{
		Right: root,
	}
	for len(stack) != 0 {
		// 取栈顶
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 先序则右子树先入栈
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		// 加入链表
		top.Left = nil
		cur.Right = top
		cur = cur.Right
	}
}

// 原地工作
func flatten3(root *TreeNode) {
	// 先序，则左子树在右子树前
	// 我们用root存放下一个待处理的节点，循环进行以下几步
	// 如果root为空，则退出循环，否则：
	// 		如果其有左子树，则将左子树放到root.Right（下一个处理的节点），并使左子树的最右节点指向原来的root.Right，然后将root.Left置空，将root后移
	// 		如果其没有左子树，则只将root后移
	for root != nil {
		if root.Left != nil {
			next := root.Left

			// root.Right放到root左子树的最下最右节点后
			newTail := root.Left
			for newTail.Right != nil {
				newTail = newTail.Right
			}
			newTail.Right = root.Right

			root.Left, root.Right = nil, next
		}
		// 后移
		root = root.Right
	}
}
