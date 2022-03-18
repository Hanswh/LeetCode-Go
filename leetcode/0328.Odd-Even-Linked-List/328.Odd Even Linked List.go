package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// c1, c2, n, n
// c1
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2 := head, head.Next
	c1, c2 := head, head.Next
	for c2 != nil && c2.Next != nil {
		c1, c2, c1.Next, c2.Next = c2.Next, c2.Next.Next, c2.Next, c2.Next.Next
	}
	c1.Next = h2
	return h1
}
