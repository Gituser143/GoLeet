Code
====

```go
func max(arr []int) int {
	maxNum := arr[0]
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func maxSlidingWindow(nums []int, k int) []int {
	sol := []int{}

	n := len(nums)

	for i := 0; i+(k-1) < n; i++ {
		maxInWindow := max(nums[i : i+k])
		sol = append(sol, maxInWindow)
	}

	return sol
}
```

Solution in mind
================

-	Solution mentioned above is a naive brute force approach. We iterate through each window and compute the max for each window.

-	We append this computed maximum to our solution.
