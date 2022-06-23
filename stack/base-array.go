package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (as *ArrayStack) IsEmpty() bool {
	if as.top <= 0 {
		return true
	}
	return false
}

func (as *ArrayStack) Push(v interface{}) {
	as.top += 1

	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

func (as *ArrayStack) Pop() interface{} {
	if as.top < 0 {
		return nil
	}

	as.top -= 1
	return as.data[as.top+1]
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
