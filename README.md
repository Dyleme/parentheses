# Parentheses

This package provides a function that determinate if the string contains correct parentheses sequence or not

- `[()]{}` - right sequence
- `[3 + 2]*(7 * {4 - 3})` - right sequence
- `[(])` - wrong sequence

## Usage
To use function you should just provide string to the function, it will return true if the sequence is correct and false if it is not.
```go
var str = "[3 + 2]*(7 * {4 - 3})"
isCorrect = balanced.IsBalanced(str)
// is correct == true
```