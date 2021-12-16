package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	left, right := dummy, head
	for n > 0 {
		right = right.Next
		n--
	}
	for right != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
