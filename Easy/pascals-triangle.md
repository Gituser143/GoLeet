Code
====

```go
func generate(numRows int) [][]int {
	sol := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		sol[i-1] = row
	}

	for i, row := range sol {
		for j, val := range row {
			if val == 0 {
				sol[i][j] = sol[i-1][j] + sol[i-1][j-1]
			}
		}
	}

	return sol
}
```

Solution in mind
================

-	Solution slice has `rows = numRows`. hence we initialise `sol` to have that many rows (int slices).

-	Each row (`int slice`) will have elements equal to the row's position. That is, ith row will have i elements in it.

-	In these i elements, the first and last are always 1. After we set this, initialisation is complete.

-	Next we iterate through elements in every row. If the element's value is 1, it stays that way, if it is 0, it's value must be computed using it's parents (element above and element above and one place left).

-	Post the iteration, the triangle will be filled.
