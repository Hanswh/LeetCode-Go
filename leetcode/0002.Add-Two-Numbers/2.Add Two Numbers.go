package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	resTail, node1, node2 := dummy, l1, l2
	carry := 0
	for node1 != nil || node2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if node1 != nil {
			v1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			v2 = node2.Val
			node2 = node2.Next
		}
		sum := v1 + v2 + carry
		resTail.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		resTail = resTail.Next
		carry = sum / 10
	}
	return dummy.Next
}
