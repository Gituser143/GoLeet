Code
====

```go
func lengthLongestPath(input string) int {
	currentPath := make([]string, 10000)

	lines := strings.Split(input, "\n")

	maxLen := 0

	for _, line := range lines {
		tabs := 0
		for _, char := range line {
			if string(char) != "\t" {
				break
			} else {
				tabs++
			}
		}
		file := strings.Trim(line, "\t")
		currentPath[tabs] = file

		if strings.Contains(file, ".") {
			temp := currentPath[:tabs+1]
			path := strings.Join(temp, "/")

			l := len(path)

			if l > maxLen {
				maxLen = l
			}
		}

	}

	return maxLen
}
```

Solution in mind
================

-	We start by iterating through line of the directory structure.

-	We keep track of the path traversed in `currentPath`. We determine at which index to place the current file/dir by determining the number of trailing tabs.

-	If we encounter a file (contains a "."), we calculate length of path upto the current index of tabs in the path. If greater than max, we update max.
