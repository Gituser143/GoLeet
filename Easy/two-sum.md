Code
====

```go
func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}
```

Solution in mind
================

-	Simple brute force.
-	Iterate from 0 to n, iterate a second number from current index + 1 to n.
-	Check if numbers at both index sum to required number.
