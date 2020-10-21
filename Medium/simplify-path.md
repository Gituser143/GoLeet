Code
====

```go
func simplifyPath(path string) string {
	// Get all directories in between slashes "/"
	dirs := strings.Split(path, "/")

	// Init Solution slice and pointer
	solutionPath := []string{}
	currentPosition := 0

	// Iterate over each directory
	for _, dir := range dirs {
		// Skip if nil or current directory
		if dir == "" {
			continue
		}

		if dir == "." {
			continue
		}

		// If previous directory, drop last element and move pointer one level back
		// If already at root (pointer = 0), do nothing
		if dir == ".." {
			if currentPosition > 0 {
				currentPosition--
				solutionPath = solutionPath[:len(solutionPath)-1]
			}
			continue
		}

		// If none of above, append directory to solution and increment pointer
		solutionPath = append(solutionPath, dir)
		currentPosition++
	}

	// Join string array with slashes and root slash
	solution := "/" + strings.Join(solutionPath, "/")
	return solution
}
```

Solution in mind
================

-	Split input path on "/" to get only directory names or relative symbols.
-	Keep track of directories in string slice. And level with pointer.
-	Iterate over above slice, disregard empty, or "." directories.
-	If previous directory ".." is encountered, move one level back. Skip if at root.
-	On encountering directory, append to solution slice and increment level pointer.
-	Return concatenated solution of solution slice, joined on "/".
