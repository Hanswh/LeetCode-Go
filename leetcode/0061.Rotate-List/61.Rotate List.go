package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转k次，每个结点的位置都可以确定，我们只需要将倒数k%size移到前面即可
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	// 移到尾部，同时求size
	size := 0
	cur := dummy
	for cur.Next != nil {
		size++
		cur = cur.Next
	}
	// 将尾结点指向头结点
	cur.Next = head
	// 重新寻找倒数第k个的前置结点
	cur = dummy
	for i := size - k%size; i > 0; i-- {
		cur = cur.Next
	}
	// 让dummy指向倒数第k个结点
	dummy.Next = cur.Next
	// 原倒数第k+1个结点变成尾结点，其Next赋空
	cur.Next = nil
	return dummy.Next
}
