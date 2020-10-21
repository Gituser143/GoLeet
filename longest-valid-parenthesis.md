Code
====

```go
func longestValidParentheses(s string) int {
	stack := []int{-1}
	l := len(s)
	valid := 0
	solution := 0

	for i := 0; i < l; i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			// ele := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				valid = i - stack[len(stack)-1]
				if valid > solution {
					solution = valid
				}
			}
		}
	}

	return solution
}
```
