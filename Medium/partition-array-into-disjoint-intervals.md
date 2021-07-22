Code
====

```go
func partitionDisjoint(nums []int) int {
	n := len(nums)

	leftMax := make([]int, n)
	rightMin := make([]int, n)

	max := nums[0]

	for i, num := range nums {
		if num > max {
			max = num
		}
		leftMax[i] = max
	}

	min := nums[n-1]

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		if num < min {
			min = num
		}
		rightMin[i] = min
	}

	for i := 1; i < n; i++ {
		if leftMax[i-1] <= rightMin[i] {
			return i
		}
	}

	return 0
}
```

Solution in mind
================

-	Solution adopted from [here](https://leetcode.com/problems/partition-array-into-disjoint-intervals/solution/)

-	We maintain a `leftMax` and `rightMin` array that holds the max value seen from the left and min value seen from the right at an index `i`.

-	The solution we require occurs when the max value seen from the left is lesser than the min value seen from the right.
