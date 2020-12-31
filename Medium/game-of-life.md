Code
====

```go
func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	indices := make(map[int][][]int, 0)
	indices[0] = [][]int{}
	indices[1] = [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] == 1 {
				count += 1
			}

			if i-1 >= 0 && board[i-1][j] == 1 {
				count += 1
			}

			if i-1 >= 0 && j+1 < n && board[i-1][j+1] == 1 {
				count += 1
			}

			if j-1 >= 0 && board[i][j-1] == 1 {
				count += 1
			}

			if j+1 < n && board[i][j+1] == 1 {
				count += 1
			}

			if i+1 < m && j-1 >= 0 && board[i+1][j-1] == 1 {
				count += 1
			}

			if i+1 < m && board[i+1][j] == 1 {
				count += 1
			}

			if i+1 < m && j+1 < n && board[i+1][j+1] == 1 {
				count += 1
			}

			if board[i][j] == 0 {
				if count == 3 {
					indices[1] = append(indices[1], []int{i, j})
				}
			} else {
				if count < 2 {
					indices[0] = append(indices[0], []int{i, j})
					continue
				}

				if count == 2 || count == 3 {
					continue
				}

				if count > 3 {
					indices[0] = append(indices[0], []int{i, j})
					continue
				}
			}
		}
	}

	for _, indice := range indices[0] {
		i, j := indice[0], indice[1]
		board[i][j] = 0
	}

	for _, indice := range indices[1] {
		i, j := indice[0], indice[1]
		board[i][j] = 1
	}
}
```

Solution in mind
================

-	We simply iterate through every element in the board and apply the rules.

-	If an element must be changed, we track it's index.

-	We later iterate through indices to be changed and update them.
