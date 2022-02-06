package leetcode

type DLinkedListNode struct {
	Key   int
	Value int
	Prev  *DLinkedListNode
	Next  *DLinkedListNode
}

func initNode(key, value int) *DLinkedListNode {
	return &DLinkedListNode{
		Key:   key,
		Value: value,
	}
}

type LRUCache struct {
	m        map[int]*DLinkedListNode
	capacity int
	head     *DLinkedListNode
	tail     *DLinkedListNode
}

func Constructor(capacity int) LRUCache {
	head := new(DLinkedListNode)
	tail := new(DLinkedListNode)
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		m:        map[int]*DLinkedListNode{},
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		// 把key对应的节点挪到双向链表头部
		this.moveNodeToHead(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		// 存在，更新值，放到链表头部
		node.Value = value
		this.moveNodeToHead(node)
	} else {
		// 不存在，加入，判断长度决定是否删除
		node = initNode(key, value)
		this.addNodeToHead(node)
		this.m[key] = node
		if len(this.m) > this.capacity {
			removed := this.removeTail()
			delete(this.m, removed.Key)
		}
	}
}

func (this *LRUCache) removeNode(node *DLinkedListNode) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (this *LRUCache) addNodeToHead(node *DLinkedListNode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) moveNodeToHead(node *DLinkedListNode) {
	this.removeNode(node)
	this.addNodeToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedListNode {
	node := this.tail.Prev
	this.removeNode(node)
	return node
}
