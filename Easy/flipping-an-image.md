Code
====

```go
func flipAndReverse(row []int) []int {
	newRow := make([]int, len(row))
	n := len(row) - 1
	for i := n; i >= 0; i-- {
		if row[n-i] == 0 {
			newRow[i] = 1
		}
	}
	return newRow
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := range A {
		A[i] = flipAndReverse(A[i])
	}

	return A
}
```

Solution in mind
================

-	Iterate through each row of image.

-	Reverse each row.

-	While reversing row, flip elements.
