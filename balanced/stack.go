package balanced

import "fmt"

type Stack struct {
	stack []int
}

func NewStack(capacity int) *Stack {
	return &Stack{stack: make([]int, 0, capacity)}
}

func (s *Stack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Push(addable int) {
	if length := len(s.stack); length == cap(s.stack) {
		temp := make([]int, length+1, 2*length+1)
		copy(temp, s.stack)
		s.stack = temp[:length]
	}

	s.stack = append(s.stack, addable)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("pop empty stack")
	}

	temp := s.stack[len(s.stack)-1]

	s.stack = s.stack[:len(s.stack)-1]

	return temp, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
