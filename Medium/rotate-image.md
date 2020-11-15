Code
====

```go
func rotate(matrix [][]int) {
	n := len(matrix)

	// Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for _, row := range matrix {
		for j := 0; j < n/2; j++ {
			row[j], row[n-j-1] = row[n-j-1], row[j]
		}
	}
}
```

Solution in mind
================

-	To rotate a matrix a simple solution would be to first Transpose the matrix and then reverse each row in the matrix.
