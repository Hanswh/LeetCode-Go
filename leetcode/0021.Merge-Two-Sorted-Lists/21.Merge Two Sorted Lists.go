package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list2 != nil && list1 != nil {
		if list2.Val < list1.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}

// 递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{}
	if list1.Val > list2.Val {
		dummy.Next = list2
		dummy.Next.Next = mergeTwoLists2(list1, list2.Next)
	} else {
		dummy.Next = list1
		dummy.Next.Next = mergeTwoLists2(list1.Next, list2)
	}

	return dummy.Next
}
