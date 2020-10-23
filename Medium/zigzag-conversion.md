Code
====

```go
func convert(s string, numRows int) string {
	n := len(s)
	numCols := 0

	if numRows > 1 {
		for n > 0 {
			if numCols%(numRows-1) == 0 {
				n = n - numRows
			} else {
				n = n - 1
			}
			numCols++
		}
	} else {
		return s
	}

	solution := make([][]byte, numRows)
	for i := range solution {
		solution[i] = make([]byte, numCols)
	}

	j := 1 // Row
	i := 0 // Col
	for x := 0; x < len(s); x++ {
		if i%(numRows-1) == 0 {
			for k := 0; k < numRows && x < len(s); k++ {
				solution[k][i] = s[x]
				x++
			}
			x--
			i++
		} else {
			solution[numRows-1-j][i] = s[x]
			j = (j + 1) % (numRows - 1)

			if j == 0 {
				j++
			}

			i++
		}
	}

	solutionString := ""

	for _, row := range solution {
		for _, letter := range row {
			if letter != 0 {
				solutionString += string(letter)
			}
		}
	}

	return solutionString
}
```

Solution in mind
================

-	First find required number of columns, done by subtracting characters occupied per column from string length.

-	Create a 2D array to store characters.

-	Keep 2 variables to keep track columns and rows. Iterate through characters in string.

-	If column is divisible by numRows-1, it can store characters in the entire column.

-	If not divisible by numRows-1, it incrementally stores values between 1 and numRows - 1.

-	Join all non empty characters. This is solution.
