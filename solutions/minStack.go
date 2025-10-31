package solutions

type Pair struct {
	val int
	min int
}

type MinStack struct {
	stack []Pair
}

func Constructor() MinStack {
	return MinStack{stack: make([]Pair, 0)}
}

func (this *MinStack) Push(val int) {
	var top Pair
	if len(this.stack) != 0 {
		top = this.stack[len(this.stack)-1]
	}
	if len(this.stack) == 0 {
		this.stack = append(this.stack, Pair{val: val, min: val})
	} else if top.min > val {
		this.stack = append(this.stack, Pair{val: val, min: val})
	} else {
		this.stack = append(this.stack, Pair{val: val, min: top.min})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
