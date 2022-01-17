package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归，函数栈
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)

	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)

	return res
}

// 手动栈
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]

		if top.Right == nil && top.Left == nil {
			res = append(res, top.Val)
			stack = stack[:len(stack)-1]
		} else {
			if top.Right != nil {
				stack = append(stack, top.Right)
				top.Right = nil
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
				top.Left = nil
			}
		}
	}

	return res
}

// Morris遍历，常数级渐近上限空间复杂度
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func postorderTraversal3(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}
