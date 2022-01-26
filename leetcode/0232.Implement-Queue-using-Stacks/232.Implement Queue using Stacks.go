package leetcode

type MyQueue struct {
	baseStack     []int
	reversedStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		baseStack:     []int{},
		reversedStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.baseStack = append(this.baseStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.reversedStack) == 0 {
		for len(this.baseStack) > 0 {
			this.reversedStack = append(this.reversedStack, this.baseStack[len(this.baseStack)-1])
			this.baseStack = this.baseStack[:len(this.baseStack)-1]
		}
	}
	val := this.reversedStack[len(this.reversedStack)-1]
	this.reversedStack = this.reversedStack[:len(this.reversedStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.reversedStack) == 0 {
		for len(this.baseStack) > 0 {
			this.reversedStack = append(this.reversedStack, this.baseStack[len(this.baseStack)-1])
			this.baseStack = this.baseStack[:len(this.baseStack)-1]
		}
	}
	val := this.reversedStack[len(this.reversedStack)-1]
	return val
}

func (this *MyQueue) Empty() bool {
	return len(this.baseStack) == 0 && len(this.reversedStack) == 0
}
