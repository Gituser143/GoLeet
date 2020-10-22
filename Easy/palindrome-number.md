Code
====

```go
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	original := x
	final := 0
	for x != 0 {
		final = final * 10
		final += x % 10
		x = x / 10
	}

	if final == original {
		return true
	}

	return false
}
```

Solution in mind
================

-	Reverse number. If reversed number == original, valid palindrome.
