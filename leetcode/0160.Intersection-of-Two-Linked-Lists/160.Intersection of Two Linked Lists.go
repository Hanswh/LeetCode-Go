package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针，空间复杂度O(1)
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headA
		} else {
			a = a.Next
		}

		if b == nil {
			b = headB
		} else {
			b = b.Next
		}
	}

	return a
}

// 哈希表
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		_, contained := m[headB]
		if contained {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
