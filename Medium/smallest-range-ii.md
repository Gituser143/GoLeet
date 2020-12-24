Code
====

```go
func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minNum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)

	N := len(A)
	max := A[N-1]
	min := A[0]
	sol := max - min

	for i := 0; i < N-1; i++ {
		a, b := A[i], A[i+1]
		sol = minNum(sol, maxNum(max-K, a+K)-minNum(min+K, b-K))
	}

	return sol
}
```

Solution in mind
================

-	Solution described [here](https://leetcode.com/problems/smallest-range-ii/solution/)
