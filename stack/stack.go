package stack

import "fmt"

type Stack struct {
	stack []int
	top   int
	size  int
}

func (s *Stack) NewStack(size int) *Stack {
	s.size = size
	s.stack = make([]int, size)
	s.top = -1
	return s
}

func (s *Stack) Push(element int) {
	if s.top == s.size-1 {
		fmt.Println("Stack overflow.")
		return
	}
	s.top += 1
	// fmt.Println(s.top)
	s.stack[s.top] = element

}

func (s *Stack) Pop() {
	if s.top == -1 {
		fmt.Println("Stack underflow.")
		return
	}
	s.stack[s.top] = 0
	s.top -= 1
}

func (s *Stack) Peek() int {
	if s.top == -1 {
		fmt.Println("Stack underflow.")
		return -1
	}
	return s.stack[s.top]
}
