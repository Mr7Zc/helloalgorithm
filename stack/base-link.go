package stack

import "fmt"

type node struct {
	next *node
	val  interface{}
}

type LinkStack struct {
	topNode *node
}

func NewLinkStack() *LinkStack {
	return &LinkStack{nil}
}

func (ls *LinkStack) IsEmpty() bool {
	return ls.topNode == nil
}

func (ls *LinkStack) Push(v interface{}) {
	ls.topNode = &node{
		next: ls.topNode,
		val:  v,
	}
}

func (ls *LinkStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	v := ls.topNode.val
	ls.topNode = ls.topNode.next
	return v
}

func (this *LinkStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
