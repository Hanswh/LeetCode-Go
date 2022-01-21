package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator interface {
	Next() int
	HasNext() bool
}

// BSTIterator1 预先构造，保存中序遍历结果
type BSTIterator1 struct {
	arr []int
}

func Constructor1(root *TreeNode) (it BSTIterator1) {
	it.inorder(root)
	return
}

func (it *BSTIterator1) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.Right)
}

func (it *BSTIterator1) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *BSTIterator1) HasNext() bool {
	return len(it.arr) > 0
}

// BSTIterator2 迭代构造
type BSTIterator2 struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{cur: root}
}

func (it *BSTIterator2) Next() int {
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}
	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator2) HasNext() bool {
	return it.cur != nil || len(it.stack) > 0
}
