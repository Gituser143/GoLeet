Code
====

```go
func validCoords(x, y, m, n int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= m || y >= n {
		return false
	}

	return true
}

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	if m == 0 {
		return mat
	}

	n := len(mat[0])

	max := m + n - 1
	base_row := m - 1
	base_col := 0

	for x := 0; x < max; x++ {
		diagonal := []int{}
		for i := 0; i <= x; i++ {
			j := x - i
			if validCoords(base_row-i, base_col+j, m, n) {
				diagonal = append(diagonal, mat[base_row-i][base_col+j])
			}
		}
		sort.Ints(diagonal)

		k := len(diagonal) - 1
		for i := 0; i <= x; i++ {
			j := x - i
			if validCoords(base_row-i, base_col+j, m, n) {
				mat[base_row-i][base_col+j] = diagonal[k]
				k--
			}
		}
	}

	return mat
}
```

Solution in mind
================

-	For an `m * n` matrix, we have `m + n - 1` diagonals. This is indicated in the variable `max`.

-	We first initialise `base_row = m - 1` (last row) and `base_col` (first column). We will populate diagonals using indices relative to these.

-	The very first diagonal will have only 1 element `matrix[base_row][base_col]`.

-	An arbitrary diagonal `x` will have elements `matrix[base_row-i][base_col+j]`, where `j` is `x-i` and `i` ranges from `0 to x-1`. A key point to note is that the values `base_row-i` and `base_col+j` are valid only if within ranges `[0, m-1]` and `[0, n-1]` respectively. Others (negatives) can be ignored.

-	Thus we iterate through diagonals with `x` and populating and sorting each of them. Upon sorting of a diagonal, we modify the matrix in-place.

-	After all diagonals have been sorted, we return the original matrix.
