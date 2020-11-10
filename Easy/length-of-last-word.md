Code
====

```go
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
```

Solution in mind
================

-	Trim all trailing and succeeding white space.

-	Split the string on space and return the length of the last word (last element).
