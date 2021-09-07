package stack

import (
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("stack: empty stack")

// Stack structure, defines such operations as:
// Pop(), Top(), Push(), Empty().
type Stack []int

// New returns empty stack of the given capacity
func New(capacity int) *Stack {
	var stack Stack = make([]int, 0, capacity)

	return &stack
}

// NewFromSlice make a stack from slice by copy
func NewFromSlice(slice []int) *Stack {
	var stack Stack = make([]int, len(slice), cap(slice))

	copy(stack, slice)

	return &stack
}

// Top returns top element of the stack.
func (s *Stack) Top() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("top: %w", ErrEmptyStack)
	}

	return (*s)[len(*s)-1], nil
}

// Push add element to the stack
// Relocate stack if capacity is end.
func (s *Stack) Push(addable int) {
	if length := len(*s); length == cap(*s) {
		temp := make([]int, length+1, 2*length+1)
		copy(temp, *s)
		*s = temp[:length]
	}

	*s = append(*s, addable)
}

// Pop delete top element of the stack.
// Return (0, error) if stack is empty.
// Return (top element, nil) if stack isn't empty.
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("pop: %w", ErrEmptyStack)
	}

	temp := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return temp, nil
}

// Empty returns true if stack is empty,
// returns false if stack isn't empty.
func (s *Stack) Empty() bool {
	return len(*s) == 0
}
