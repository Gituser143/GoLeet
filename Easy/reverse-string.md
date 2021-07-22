Code
====

```go
func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}
```

Solution in mind
================

-	Swap front and back elements, while moving towards the centre.
