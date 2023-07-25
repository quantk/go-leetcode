package main

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1
		value := (*s)[index]
		*s = (*s)[:index]
		return value
	}
}

type MyQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: &Stack{},
		s2: &Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) amortize() {
	for this.s1.IsEmpty() == false {
		this.s2.Push(this.s1.Pop())
	}
}

func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() == true {
		this.amortize()
	}

	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() == true {
		this.amortize()
	}

	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}
