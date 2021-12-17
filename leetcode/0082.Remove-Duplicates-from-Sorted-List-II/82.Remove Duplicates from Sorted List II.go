package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 因为head可能被去除，使用dummy
	dummy := &ListNode{0, head}
	// cur表示去除重复结点后链表的尾结点，初始指向dummy表示还没有
	cur := dummy
	// 过程是遍历寻找不重复的结点————对于当前的cur，判断cur.Next是否可以加入
	// 当前cur.Next如果和cur.Next.Next值不同，则可以加入，我们将cur后移，让原来的cur.Next加入
	// 当前cur.Next如果和cur.Next.Next值相同，则接下来所有这个值的结点都不行，我们后移cur.Next直到找到值不同的位置，再重新判断
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
