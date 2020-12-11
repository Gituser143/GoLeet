Code
====

```go
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	i := 0
	for _, n := range nums {
		if i < 2 || n > nums[i-2] {
			nums[i] = n
			i += 1
		}
	}

	return i
}
```

Solution in mind
================

-	Keep a pointer to index `i`. Iterate through array `nums`, using each element `num` in `nums`.

-	Since a number can occur at most twice, check if the current number is greater than the number at `i-2`.

-	If it is, place the number at `i` and move `i` forward by 1 place.

-	If not, iterate to the next number and check if can be placed at `i` by checking if it is greater than the number at `i-2`.
