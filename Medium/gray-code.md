Code
====

```go
func grayCode(n int) []int {
	sol := []int{0}

	for i := 1; i <= n; i++ {
		prevLength := len(sol)
		mask := 1 << (i - 1)

		for j := prevLength - 1; j >= 0; j-- {
			sol = append(sol, mask+sol[j])
		}
	}

	return sol
}
```

Solution in mind
================

Solution adopted from [here](https://leetcode.com/problems/gray-code/solution/)
