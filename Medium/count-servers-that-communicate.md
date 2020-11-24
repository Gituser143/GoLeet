Code
====

```go
func sum(arr []int) (int, int) {
	s := 0
	lastIndex := -1
	for i, num := range arr {
		if num > 0 {
			lastIndex = i
		}
		s += num
	}
	return s, lastIndex
}

func countServers(grid [][]int) int {
	sol := 0
	colSums := make(map[int]int, 0)

	for _, row := range grid {
		rowSum, index := sum(row)

		if rowSum > 1 {
			sol += rowSum
			continue

		} else if rowSum == 1 {
			colSum, ok := colSums[index]
			if !ok {
				colSum = 0
				for _, row := range grid {
					colSum += row[index]
				}

				colSums[index] = colSum
			}

			if colSum > 1 {
				sol += 1
			}
		}
	}
	return sol
}
```

Solution in mind
================

-	To get the count of servers, we iterate through each row of the grid.

-	We compute the sum of each row, if this sum is greater than 1, there are multiple servers in the row which can communicate with each other. We add this sum to the solution.

-	If the sum is 1, this means there is one server on the grid and we need to check if there are any servers in that column. If there are, we can add the single server on the row to the solution.

-	To check if there are multiple servers on the same column, we make use of a map, which keeps track of column sums. This is particularly helpful so that we need not re compute column sums.

-	If the column sum exists for the required column and is greater than 1, we increment the solution count by 1.

-	If the column sum does not exist in map, we compute it by iterating over cells in the column and then add it to the map, following which we check to see if the sum is greater than 1. If yes, we increment solution count by 1.
