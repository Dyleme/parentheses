package balanced

import "fmt"

func IsBalanced(str string) bool {
	stack := make([]int, 0)
	fmt.Println(stack)
	for _, char := range str {
		switch char {
		case '{':
			stack = append(stack, 1)
		case '(':
			stack = append(stack, 2)
		case '[':
			stack = append(stack, 3)
		case '}':
			if len(stack) != 0 &&  stack[len(stack)-1] == 1 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ')':
			if len(stack) != 0 && stack[len(stack)-1] == 2 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) != 0 && stack[len(stack)-1] == 3 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		fmt.Println(stack)
	}
	return len(stack) == 0
}
