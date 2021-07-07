Code
====

```go
func kthSmallest(matrix [][]int, k int) int {
	singleRow := []int{}
	for _, row := range matrix {
		singleRow = append(singleRow, row...)
	}

	sort.Ints(singleRow)
	return singleRow[k-1]
}
```

Solution in mind
================

-	Flatten the matrix into a single row.

-	Sort the single row and return the kth element.
