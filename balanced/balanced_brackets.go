package balanced

import "fmt"

type Stack struct {
	stack []int
	len int
	cap int
}

func NewStack(cap int) *Stack {
	return &Stack{make([]int, cap),0,cap}
}

func (s *Stack) top() int {
	return s.stack[s.len - 1]
}

func (s *Stack) push(addble int)  {
	if s.len == s.cap - 1 {
		temp := make([]int, s.cap + 2)
		copy(temp, s.stack)
		s.stack = temp
	}
	s.stack[s.len] = addble
	s.len++
}

func (s *Stack) pop() (int, error) {
	if s.len == 0 {
		return 0, fmt.Errorf("pop empty stack")
	}
	s.len--
	return s.stack[s.len], nil
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func IsBalanced(str string) bool {
	stack := NewStack(len(str)/2 + 1)
	var err error
	for _, char := range str {
		switch char {
		case '{':
			stack.push(1)
		case '(':
			stack.push(2)
		case '[':
			stack.push(3)
		case '}':
			if !stack.IsEmpty() &&  stack.top() == 1 {
				_, err = stack.pop()
			} else {
				return false
			}
		case ')':
			if !stack.IsEmpty() && stack.top() == 2 {
				_, err = stack.pop()
			} else {
				return false
			}
		case ']':
			if !stack.IsEmpty() && stack.top() == 3 {
				_, err = stack.pop()
				} else {
				return false
			}
		}
		if err != nil {
			return false
		}
	}
	return stack.IsEmpty()
}
