package leetcode

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

func (this *MyStack) Push(x int) {
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1 = append(this.queue1, x)
	for len(this.queue2) > 0 {
		this.queue1 = append(this.queue1, this.queue2[0])
		this.queue2 = this.queue2[1:]
	}
}

func (this *MyStack) Pop() (val int) {
	val = this.queue1[0]
	this.queue1 = this.queue1[1:]
	return
}

func (this *MyStack) Top() (val int) {
	val = this.queue1[0]
	return
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
