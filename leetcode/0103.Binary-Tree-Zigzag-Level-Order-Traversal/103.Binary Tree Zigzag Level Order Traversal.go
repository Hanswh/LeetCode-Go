package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var level []int
	d := false // 本层方向，true表示从左到右，false表示反向
	dq := []*TreeNode{root}
	count := 1
	for count > 0 {
		if d {
			// 从左到右
			top := dq[0]
			dq = dq[1:]
			level = append(level, top.Val)
			if top.Left != nil {
				dq = append(dq, top.Left)
			}
			if top.Right != nil {
				dq = append(dq, top.Right)
			}
			count--
		} else {
			// 从右往左
			top := dq[len(dq)-1]
			dq = dq[:len(dq)-1]
			level = append(level, top.Val)
			if top.Right != nil {
				dq = append([]*TreeNode{top.Right}, dq...)
			}
			if top.Left != nil {
				dq = append([]*TreeNode{top.Left}, dq...)
			}
			count--
		}
		if count == 0 {
			// 本层遍历结束，加入结果集，进行下一层
			res = append(res, level)
			d, level, count = !d, []int{}, len(dq)
		}
	}
	return res
}
