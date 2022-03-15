package leetcode

type Node struct {
	val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	levelCount := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		levelCount--
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		if levelCount == 0 {
			cur.Next = nil
			levelCount = len(queue)
		} else {
			cur.Next = queue[0]
		}
	}
	return root
}
