Code
====

```go
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func createSortedArray(instructions []int) int {
	mod := 1000000007
	nums := make([]int, len(instructions))
	cost := 0

	for length, num := range instructions {
		l, r := sort.Search(length, func(i int) bool { return nums[i] >= num }), sort.Search(length, func(i int) bool { return nums[i] > num })

		cost += min(l, length-r)

		copy(nums[r+1:length+1], nums[r:length+1])
		nums[r] = num
	}
	return cost % mod
}
```

Solution in mind
================

-	Find the left and right most indices where the element can be inserted.

-	Check which has lesser cost and increment cost.

-	Insert element into `nums`. Insertion happens, by shifting the elements after the rightmost position by one space, and inserting the element into that newly created space.
