package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表
func detectCycle1(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if m[cur] {
			return cur
		}
		m[cur] = true
		cur = cur.Next
	}
	return nil
}

// 双指针
// 假设无环部分长a，快慢指针相遇时，慢指针在环中走了b，在环中还有长度c没走，放在坐标系中可以得到a = c
func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
