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
type Solution1 struct {
	arr []int
}

func Constructor1(head *ListNode) Solution1 {
	solution := Solution1{arr: []int{}}
	for head != nil {
		solution.arr = append(solution.arr, head.Val)
		head = head.Next
	}
	return solution
}

func (s *Solution1) GetRandom() int {
	return s.arr[rand.Int()%len(s.arr)]
}

// Solution2 存头结点
type Solution2 struct {
	head *ListNode
	len  int
}

func Constructor2(head *ListNode) Solution2 {
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	return Solution2{head: head, len: count}
}

func (s *Solution2) GetRandom() int {
	index := rand.Int() % s.len
	cur := s.head
	for index > 0 {
		cur = cur.Next
		index--
	}
	return cur.Val
}
