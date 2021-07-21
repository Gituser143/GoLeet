Code
====

```go
func pushDominoes(dominoes string) string {
	if len(dominoes) == 0 || len(dominoes) == 1 {
		return dominoes
	}

	tosee := "L" + dominoes + "R"
	round := []byte(tosee)

	left := 0
	right := 1

	// while right does not reach the end
	for right < len(round) {
		// move right till a character is seen
		for round[right] == '.' {
			right++
		}

		switch {
		// if [left:right] is L...L, we fill with L
		case round[left] == 'L' && round[right] == 'L':
			for i := left + 1; i < right; i++ {
				round[i] = 'L'
			}

		// if [left:right] is R..R, we fill with R
		case round[left] == 'R' && round[right] == 'R':
			for i := left + 1; i < right; i++ {
				round[i] = 'R'
			}

		// if [left:right] is R...L, we fill up to center with R and from
		// center to right with L. If a center exists (left-right) is odd, it stays as '.'
		case round[left] == 'R' && round[right] == 'L':
			for i := left + 1; i <= (left+right-1)/2; i++ {
				round[i] = 'R'
			}
			for i := right - ((left+right-1)/2 - left); i < right; i++ {
				round[i] = 'L'
			}
		}

		// move left pointer upto processed right
		left = right
		right++
	}

	return string(round[1 : len(round)-1])
}
```
