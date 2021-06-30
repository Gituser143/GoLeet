Code
====

```go
func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		} else {
			count += 1
		}
	}

	if count > maxCount {
		maxCount = count
	}

	return maxCount
}
```

Solution in mind
================

-	Keep a count of how many `1's` have been encountered.

-	If a 0 is encountered, check if count is largest seen and update `maxCount`. Reset the count to 0.
