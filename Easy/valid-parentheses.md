Code
====

```go
func isValid(s string) bool {
	bracketMap := make(map[string]string, 0)
	bracketMap[")"] = "("
	bracketMap["]"] = "["
	bracketMap["}"] = "{"

	stack := []string{}

	for _, val := range s {
		letter := string(val)
		if _, ok := bracketMap[letter]; !ok {
			stack = append(stack, letter)
		} else {
			if len(stack) == 0 {
				return false
			}
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != bracketMap[letter] {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
```

Solution in mind
================

-	Iterate through string, if opening bracket occurs, append to stack.

-	If closing bracket occurs, pop the top most element from stack and if popped element is not the matching opening bracket, string is invalid. If the stack is empty, there are excess closing brackets and the string is invalid.

-	At the end of string iteration if any brackets are left, there are excess open brackets and string is thus invalid.
