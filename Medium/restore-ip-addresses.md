Code
====

```go
func dfs(s string, blocks int, path string, result []string) []string {
	// If input string is greater than 12, no possible solution
	if len(s) > 12 {
		return result
	}

	// If the current string has more than 4 blocks of digits, invalid so return
	if blocks > 4 {
		return result
	}

	// If no more string is left to parse and there are 4 blocks in the IP, append to result
	if blocks == 4 && len(s) == 0 {
		result = append(result, path[:len(path)-1])
		return result
	}

	// Try t create IPs block wise
	// Block width can be 1-3 [0-255]
	for i := 1; i < 4; i++ {
		// Avoid index out of range issues
		if i > len(s) {
			continue
		}

		// Get integer value of block
		num, err := strconv.Atoi(s[:i])
		if err != nil {
			return result
		}

		// Make sure block is either 0 or a number less than 256 but does not start with 0
		if string(s[:i]) == "0" || (string(s[0]) != "0" && num < 256) {
			// Recursively call to see if further blocks can fit valid IP
			// valid IP is stored as path, remaining string is parsed to fit IP
			result = dfs(s[i:], blocks+1, path+string(s[:i])+".", result)
		}
	}
	return result
}

func restoreIpAddresses(s string) []string {
	result := []string{}

	result = dfs(s, 0, "", result)
	return result

}
```

Solution in mind
================

-	To generate IPs, Try to generate block by block.
-	Iterate through block sizes of 1-3, strip block from original string and add block to path.
-	Ensure block has number between 0-255 and does not have trailing 0s.
-	Recursively recreate blocks on new stripped string, while string is not empty.
-	If string is empty and a valid path is found (4 blocks) append to results list.
