Code
====

```go
func customSortString(order string, str string) string {
	positionMap := make(map[rune]int)
	for i, c := range order {
		positionMap[c] = i
	}

	bs := []rune(str)

	sort.Slice(bs, func(i, j int) bool {
		x, ok := positionMap[bs[i]]
		if !ok {
			x = i
		}

		y, ok := positionMap[bs[j]]
		if !ok {
			y = j
		}

		return x < y

	})

	return string(bs)
}
```

Solution in mind
================

-	We maintain a position map (`positionMap`), which holds the index of a given letter.

-	We then perform a straight forward sort, but in the less function, we compare the positions of the two characters by their index in `positionMap` to determine if less.
