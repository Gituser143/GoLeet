Code
====

```go
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := false

	if s != "" && (p[0] == s[0] || string(p[0]) == ".") {
		firstMatch = true
	}

	if len(p) >= 2 && string(p[1]) == "*" {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
```

Solution
========

Solution on [leetcode](https://leetcode.com/problems/regular-expression-matching/solution/)
