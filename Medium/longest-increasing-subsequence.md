Code
====

```go
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0

	for i, num := range nums {
		dp[i] = 1
		if i > 0 {
			for j := 0; j < i; j++ {
				if num > nums[j] {
					dp[i] = max(dp[i], dp[j]+1)
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
```

Solution in mind
================

-	We approach this problem as a dynamic programming problem. We maintain an array (`dp`) of length `n == len(nums)`.

-	The value of the ith element in `dp` is the number of increasing subsequences that can be formed upto that index in `nums` (upto `nums[i]`).

-	We have the base case that every element in itself can be a subsequence, thus every element in the array is initialised to 1 (`dp[i] = 1`).

-	At an index `i`, we iterate from 0 upto `i-1` and check if any element is lesser than the current number at `i`. If yes, `dp[i]` will be updated to `dp[j] + 1`, provided `dp[j] + 1` is greater than `dp[i]`.

-	Our solution will be the maximum element of the `dp` array.
