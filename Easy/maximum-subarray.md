Code
====

```go
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
```

Solution in mind
================

-	We iterate through the array and replace the number at a given position by the maximum sum that can be created till that position.

-	The number at a given position is thus set to the `max(nums[i], nums[i]+nums[i-1])`.

-	Thus if a sum is lesser than the number itself, we're better off by using just the number.
