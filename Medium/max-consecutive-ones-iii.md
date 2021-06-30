Code
====

```go
func longestOnes(nums []int, k int) int {
	maxNum := 0
	zeros := 0
	i := 0

	// Iterate using i as startIdx and j as endIdx
	// j-i+1 is now the window
	for j, num := range nums {
		// If a 0 is encountered, increment the zero count
		if num == 0 {
			zeros += 1
		}

		// If more than k 0s have been encountered,
		// keep moving start till a 0 is dropped from
		// the beginning
		for zeros > k {
			if nums[i] == 0 {
				zeros -= 1
			}
			i += 1
		}

		// Calculate window size as end - beginning + 1
		l := j - i + 1
		if l > maxNum {
			maxNum = l
		}
	}

	return maxNum
}
```
