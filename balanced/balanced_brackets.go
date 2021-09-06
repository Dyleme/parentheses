package balanced

func IsBalanced(str string) bool {
	stack := NewStack(len(str)/2 + 1)

	var err error

	for _, char := range str {
		switch char {
		case '{':
			stack.Push(int('{'))
		case '(':
			stack.Push(int('('))
		case '[':
			stack.Push(int('['))
		case '}':
			if !stack.IsEmpty() && stack.Top() == int('{') {
				_, err = stack.Pop()
			} else {
				return false
			}
		case ')':
			if !stack.IsEmpty() && stack.Top() == int('(') {
				_, err = stack.Pop()
			} else {
				return false
			}
		case ']':
			if !stack.IsEmpty() && stack.Top() == int('[') {
				_, err = stack.Pop()
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
