Code
====

```go
func decodeAtIndex(S string, K int) string {
	size := 0
	n := len(S)

	for _, letter := range S {
		if unicode.IsDigit(letter) {
			num, _ := strconv.Atoi(string(letter))
			size *= num
		} else {
			size += 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		letter := S[i]
		K = K % size

		if K == 0 && unicode.IsLetter(rune(letter)) {
			return string(letter)
		}

		if unicode.IsDigit(rune(letter)) {
			num, _ := strconv.Atoi(string(letter))
			size = size / num
		} else {
			size -= 1
		}
	}

	return ""
}
```

Solution
========

-	Solution stated [here](https://leetcode.com/problems/decoded-string-at-index/solution/)
