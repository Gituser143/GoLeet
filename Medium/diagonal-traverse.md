Code
====

```go
func findDiagonalOrder(matrix [][]int) []int {
	sol := []int{}

	m := len(matrix)
	if m == 0 {
		return sol
	}

	n := len(matrix[0])

	max := m + n - 1

	for x := 0; x < max; x++ {
		if x%2 == 0 {
			for i := 0; i <= x; i++ {
				if i >= n || x-i >= m {
					continue
				}
				sol = append(sol, matrix[x-i][i])
			}
		} else {
			for i := 0; i <= x; i++ {
				if i >= m || x-i >= n {
					continue
				}
				sol = append(sol, matrix[i][x-i])
			}
		}
	}

	return sol
}
```

Solution in mind
================

-	For a given `M * N` matrix, we have `M + N - 1` diagonals to traverse.

-	In each of these diagonals, the row number `i` and the column number `j` sum upto the same value `x`, where `x` ranges from `0 to M + N - 2`.

-	Thus we iterate through each diagonal, (0, to M+N-2) using `x`, and take elements from `i`, `x-i`, such that both these indices are within column/row limits.

-	Furthermore, as even diagonals are reversed, we just need to append elements with indices `x-i,`, `i` instead of `i`, `x-i`. Thus with `x` iterations, we are able to traverse and track elements through each diagonal.
