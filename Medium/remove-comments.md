Code
====

```go
func removeComments(source []string) []string {
	inBlock := false
	sol := []string{}
	newLine := ""
	for _, line := range source {
		i := 0

		if !inBlock {
			newLine = ""
		}

		for i < len(line) {
			if i < len(line)-1 && string(line[i:i+2]) == "/*" && !inBlock {
				inBlock = true
				i++
			} else if i < len(line)-1 && string(line[i:i+2]) == "*/" && inBlock {
				inBlock = false
				i++
			} else if i < len(line)-1 && !inBlock && string(line[i:i+2]) == "//" {
				break
			} else if !inBlock {
				newLine += string(line[i])
			}
			i++
		}

		if newLine != "" && !inBlock {
			sol = append(sol, newLine)
		}
	}

	return sol
}
```

Solution in mind
================

We need to parse the `source` line by line. Our state is that we either are in a block comment or not.

-	If we start a block comment and we aren't in a block, then we will skip over the next two characters and change our state to be in a block.

-	If we end a block comment and we are in a block, then we will skip over the next two characters and change our state to be not in a block.

-	If we start a line comment and we aren't in a block, then we will ignore the rest of the line.

-	If we aren't in a block comment (and it wasn't the start of a comment), we will record the character we are at.

-	At the end of each line, if we aren't in a block, we will record the line.
