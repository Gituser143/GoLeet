Code
====

```go
func lengthOfLongestSubstring(s string) int {
	solution := 0
	l := len(s)
	for i := 0; i < l; i++ {
		var substring int
		letters := make(map[string]bool, 0)

		letters[string(s[i])] = true
		substring += 1

		for j := i + 1; j < l; j++ {
			if _, ok := letters[string(s[j])]; ok {
				break
			} else {
				substring += 1
				letters[string(s[j])] = true
			}
		}

		if substring > solution {
			solution = substring
		}
	}
	return solution
}
```

Solution in mind
================

-	Iterate over every character.
-	Go through every character from the previous character and increment count for every unique character.
-	Break on occurrence of non unique char. If current substring length greater than, previous, update solution with current.
