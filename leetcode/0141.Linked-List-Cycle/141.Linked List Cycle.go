package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针，两指针相遇则有环
func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 哈希表优化蛮力搜索
func hasCycle2(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, contained := m[head]; contained {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}
