Code
====

```go
func removeDuplicates(s string) string {
	stack := []rune{}

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
		} else {
			if c == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}

	return string(stack)
}
```

Solution in mind
================

-	We iterate through the given string, character by character while maintaining a stack.

-	A character is pushed onto the stack only if:

	-	The stack is empty
	-	The character on top of the stack is not the same as the current character

-	If the character on top is the same as the current character, we pop it from the stack as it is a duplicate.

-	After iterating through the string, the stack now has the required new string, character by character.
