Code
====

```go
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])

	if m*n != r*c {
		return mat
	}

	singleRow := []int{}
	for _, row := range mat {
		singleRow = append(singleRow, row...)
	}

	newMat := make([][]int, r)
	for i := range newMat {
		row := singleRow[:c]
		singleRow = singleRow[c:]
		newMat[i] = row
	}

	return newMat
}
```

Solution in mind
================

-	A solution is legal only if the product of the number of rows and columns is equal to the product of required rows and columns (`m * n == r * c`).

-	A single row (`singleRow`) is first populated to hold all the elements.

-	A new matrix (`newMat`) is created with the required number of rows and each row is populated by removing `c` number of elements from `singleRow` and putting it into the row of the new matrix.
