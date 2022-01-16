package leetcode

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution interface {
	GetRandom() int
}

// Solution1 数组缓存，可以随机访问
type Solution1 []int

func Constructor1(head *ListNode) (s Solution1) {
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	return
}

func (s Solution1) GetRandom() int {
	return s[rand.Intn(len(s))]
}

// Solution2 存头结点
type Solution2 struct {
	head *ListNode
	len  int
}

func Constructor2(head *ListNode) (s Solution2) {
	s.head = head
	s.len = 0
	for head != nil {
		s.len++
		head = head.Next
	}
	return
}

func (s Solution2) GetRandom() int {
	index := rand.Intn(s.len)
	cur := s.head
	for index > 0 {
		cur = cur.Next
		index--
	}
	return cur.Val
}
