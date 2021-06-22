Code
====

```go
func getRow(rowIndex int) []int {
	prevRow := []int{}
	for i := 0; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j, val := range row {
			if val != 1 {
				row[j] = prevRow[j] + prevRow[j-1]
			}
		}

		prevRow = row
		if i == rowIndex {
			return row
		}
	}

	return []int{}
}
```

Solution in mind
================

-	Populate triangle, row by row. We keep track of previous row (`prevRow`) and the current row (`row`) is created by summing values from the previous.

-	In every row, the first and last element are `1`. Thus we can distinguish elements to be calculated as those that do not have their value equal to `1`.

-	A value at index `j` in a row can be calculated as the sum of `j` and `j-1` values of the previous row.

-	As we populate these rows, we return the row created when `i == rowIndex`.
