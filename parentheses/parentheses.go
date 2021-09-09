package parentheses

import (
	"strings"
)

// IsBalanced function check if provided string contains correct sequence of brackets.
// Returns true if to contains, false if it is not.
func IsBalanced(str string) bool {
	stack := make([]rune, 0, len(str))
	pairBracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range str {
		if strings.ContainsRune("{([", char) {
			stack = append(stack, char)
		} else if strings.ContainsRune(")}]", char) {
			if len(stack) != 0 && stack[len(stack)-1] == pairBracket[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
