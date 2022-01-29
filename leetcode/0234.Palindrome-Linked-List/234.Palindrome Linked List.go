package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 放到数组
func isPalindrome1(head *ListNode) bool {
	vals := make([]int, 0)
	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}

	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-i-1] {
			return false
		}
	}
	return true
}

// 原地工作
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

// 双指针反转列表
func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

// 快慢指针找前半部分的尾结点
func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
