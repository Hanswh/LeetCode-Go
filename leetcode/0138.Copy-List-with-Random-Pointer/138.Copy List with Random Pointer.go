package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 借助哈希表存储原链表每个节点的位置，先构建链表，再依次根据原链表的随机指针指向的节点的位置获取新链表随机指针指向的节点，构建。
func copyRandomList1(head *Node) *Node {
	idx := make(map[*Node]int)
	position := 0

	dummy := &Node{}
	n1 := head
	n2 := dummy
	for n1 != nil {
		idx[n1] = position
		position++

		n2.Next = &Node{Val: n1.Val}
		n2 = n2.Next
		n1 = n1.Next
	}

	n1 = head
	n2 = dummy.Next
	for n1 != nil {
		if n1.Random != nil {
			n2.Random = getN(dummy.Next, idx[n1.Random])
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return dummy.Next
}

// 获取第n个节点
func getN(head *Node, idx int) *Node {
	for idx > 0 {
		head = head.Next
		idx--
	}
	return head
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	tempHead := copyNodeToLinkedList(head)
	return splitLinkedList(tempHead)
}

func splitLinkedList(head *Node) *Node {
	cur := head
	head = head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}
	return head
}

func copyNodeToLinkedList(head *Node) *Node {
	cur := head
	for cur != nil {
		node := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next, cur = node, cur.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	return head
}
