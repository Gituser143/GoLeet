Code
====

```go
func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
```

Solution in mind
================

Intuition
---------

Since this question is asking us to remove all elements of the given value in-place, we have to handle it with `O(1)` extra space. How to solve it? We can keep two pointers `i` and `j`, where iii is the slow-runner while `j` is the fast-runner.

Algorithm
---------

When `nums[j]` equals to the given value, skip this element by incrementing `j`. As long as `nums[j] != val`, we copy `nums[j]` to `nums[i]` and increment both indexes at the same time. Repeat the process until `j` reaches the end of the array and the new length is `i`.
