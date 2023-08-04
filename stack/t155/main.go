package main

/**
 * 最小栈, 用两个栈实现, 一个栈正常存储, 另一个栈存储最小值
 * 关键是要理解栈的特性，只有一个出口，最小栈维护当前对应位置的最小值方便在 pop 时同步更新
 */

// MinStack ...
type MinStack struct {
	stack []int
	ms    []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		ms:    make([]int, 0),
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.ms) == 0 {
		s.ms = append(s.ms, val)
	} else {
		s.ms = append(s.ms, min(s.ms[len(s.ms)-1], val))
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[0 : len(s.stack)-1]
	s.ms = s.ms[0 : len(s.ms)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.ms[len(s.ms)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	println(obj.GetMin())
	obj.Pop()
	println(obj.Top())
	println(obj.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
