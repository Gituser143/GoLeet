Code
====

```go
func isPalindrome(s string) bool {
	l := len(s)

	if l == 1 {
		return true
	}

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var solution string
	var solutionLen int

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && solutionLen < (j-i) {
				solution = s[i:j]
				solutionLen = j - i
			}
		}
	}

	return solution
}
```

Solution in mind
================

-	Brute force approach.
-	Iterate through characters of string.
-	For each iteration, check if sting from current character to end is a palindrome. End is initialised to next char.
-	If yes, return palindrome, if not, increment end by one character until it reaches end and repeat.
