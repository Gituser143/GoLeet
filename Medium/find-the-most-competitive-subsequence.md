Code
====

```go
func mostCompetitive(nums []int, k int) []int {
	sol := []int{}
	n_nums := len(nums)

	for i, num := range nums {
		n := len(sol)
		for n > 0 && num < sol[n-1] && n_nums-i+n-1 >= k {
			sol = sol[:n-1]
			n -= 1
		}

		if n < k {
			sol = append(sol, num)
		}

	}

	return sol
}
```

Solution in mind
================

-	Solution based off discussion post, given [here](https://leetcode.com/problems/find-the-most-competitive-subsequence/discuss/952786/JavaC++Python-One-Pass-Stack-Solution)
