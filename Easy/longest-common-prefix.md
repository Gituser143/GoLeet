Code
====

```go
func longestCommonPrefix(strs []string) string {
	numWords := len(strs)

	if numWords == 1 {
		return strs[0]
	} else if numWords == 0 {
		return ""
	}

	i := 0

	for {
		if len(strs[0]) <= i {
			return strs[0][:i]
		}
		char := strs[0][i]
		for j := 1; j < numWords; j++ {
			if len(strs[j]) <= i {
				return strs[0][:i]
			}
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
		i++
	}
}
```

Solution in mind
================

-	Iterate through indices of each string. If index exceeds or equals length of string, return string upto index.

-	Compare character at ith index of every string with that of 0th string. If a string's character doesn't match, return string upto i.
