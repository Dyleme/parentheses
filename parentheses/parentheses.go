package parentheses

import (
	"math/rand"
	"strings"
	"time"
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

// GenerateBrackets generate random sequence of brackets provided length.
// Argument length must be positive number, otherwise function panics.
func GenerateBrackets(length int) string {
	if length < 1 {
		panic("length of sequences can not be zero or negative number")
	}

	brackets := "(){}[]"

	rand.Seed(time.Now().Unix())

	var sb strings.Builder

	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(brackets[rand.Intn(len(brackets))])) //nolint:gosec //Should be used according to  the task.
	}

	return sb.String()
}
