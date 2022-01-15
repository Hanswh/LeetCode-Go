package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = tmp
	}
	return dummy.Next
}
