Code
====

```go
func isValidSudoku(board [][]byte) bool {
	boxes := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		boxes[i] = make(map[byte]bool, 0)
		rows[i] = make(map[byte]bool, 0)
		cols[i] = make(map[byte]bool, 0)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 46 {
				continue
			}

			if _, ok := rows[i][board[i][j]]; ok {
				return false
			} else {
				rows[i][board[i][j]] = true
			}

			if _, ok := cols[j][board[i][j]]; ok {
				return false
			} else {
				cols[j][board[i][j]] = true
			}

			boxnum := int(i/3)*3 + int(j/3)

			if _, ok := boxes[boxnum][board[i][j]]; ok {
				return false
			} else {
				boxes[boxnum][board[i][j]] = true
			}
		}
	}

	return true
}
```

Solution in mind
================

-	To validate the sudoku, all we have to ensure is that no element occurs twice in a row, column or sub box.

-	To ensure this, we keep an array of sets, each containing a row, column or sub box.

-	To get the index of box an element at position i, j belongs to, we use the formula `boxnum := int(i/3)*3 + int(j/3)`. Which is the code equivalent of `box_index = (row / 3) * 3 + col / 3`, where / is integer division.

-	We iterate through each cell of the sudoku and check if that element has occurred previously in the same row, column or sub box, if yes, we return false.

-	If no element occurs twice, we return true.
