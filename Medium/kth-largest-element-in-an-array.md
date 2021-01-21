Code
====

```go
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
```

Solution in mind
================

-	Sort given array and return the kth element from the end.
