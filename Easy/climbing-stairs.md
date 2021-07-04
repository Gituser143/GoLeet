Code
====

```go
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
```

Solution in mind
================

-	To get the number of ways to reach a step `n`, we add the number of ways to get to `n-1` and `n-2`. This is because there are only 2 ways to get to `n` in one step, from either `n-1` or `n-2`.

-	We thus populate the `dp` array iteratively using the sum of previous two steps to get the number of ways we can reach `n`.
