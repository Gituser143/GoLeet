Code
====

```go
func sortedSquares(nums []int) []int {
	sol := []int{}

	for _, num := range nums {
		sol = append(sol, num*num)
	}

	sort.Ints(sol)

	return sol
}
```

Solution in mind
================

-	Iterate through array keeping track of squares.

-	Sort array holding squares
