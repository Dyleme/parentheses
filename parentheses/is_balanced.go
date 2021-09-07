package balanced

import (
	"parentheses/stack"

	"strings"
)

// IsBalanced function check if provided string contains correct sequence of brackets.
func IsBalanced(str string) bool {
	s := stack.New(len(str)/2 + 1)

	for _, char := range str {
		if strings.ContainsRune("{([", char) {
			s.Push(int(char))
		} else if strings.ContainsRune(")}]", char) {
			if ok := tryPop(s, pairBracket(char)); !ok {
				return false
			}
		}
	}

	return s.Empty()
}

// tryPop function get stack and rune that it tries to pop.
// If at top of stack is provided rune, than function pop it and returns true.
// If stack is empty or top of it the stack isn't provided rune, returns false.
func tryPop(s *stack.Stack, char rune) bool {
	if s.Empty() {
		return false
	}

	if s.Top() == int(char) {
		_, err := s.Pop()
		if err != nil {
			return false
		}
	}

	return true
}

// pairBracket returns pair bracket to any bracket.
// Returns ' ' (space rune) if provided rune isn't bracket.
func pairBracket(bracket rune) rune {
	switch bracket {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	}

	return ' '
}
