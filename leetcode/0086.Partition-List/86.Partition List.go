package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	// 让left指向第一个比x值大的结点的前一个结点
	// 接下来让right指向第一个比x值大的结点，并依次向后遍历，如果存在比x值小的结点，将其尾插至left后
	// 因为要更新，我们对right.Next进行判断
	left := dummy
	for left.Next != nil && left.Next.Val < x {
		left = left.Next
	}
	// 已经到达链表尾部，说明已经满足要求
	if left.Next == nil {
		return head
	}
	right := left.Next
	for right.Next != nil {
		if right.Next.Val < x {
			// 将right.Next移至left后
			tmp := right.Next.Next
			right.Next.Next = left.Next
			left.Next = right.Next
			right.Next = tmp
			// 后移left，保持其为小值链表的尾部
			left = left.Next
		} else {
			right = right.Next
		}
	}
	return dummy.Next
}
