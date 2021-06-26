Code
====

```go
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func evalRPN(tokens []string) int {

	evalMap := make(map[string]func(a, b int) int)
	evalMap["+"] = add
	evalMap["-"] = subtract
	evalMap["*"] = multiply
	evalMap["/"] = divide

	stack := []int{}

	for _, token := range tokens {
		if f, ok := evalMap[token]; ok {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			x := f(a, b)

			stack = append(stack, x)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
```

Solution in mind
================

-	We maintain a stack to keep track of tokens and perform operations. The stack only holds integers in it.

-	We iterate through tokens and if a number is encountered, we push it onto the stack.

-	In the case where an operation is encountered, we pop off the last two elements and calculate the result of the operation.

-	The result of the operation is appended back onto the stack.

-	The end result will be the sole element in the stack which can be returned as solution.
