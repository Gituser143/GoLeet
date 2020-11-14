Code
====

```go
func fourSum(nums []int, target int) [][]int {
	solution := [][]int{}

	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n-3; i++ {
		for i != 0 && nums[i] == nums[i-1] && i < n-3 {
			i++
		}
		for j := i + 1; j < n-2; j++ {
			for j != i+1 && nums[j] == nums[j-1] && j < n-2 {
				j++
			}

			lb := j + 1
			ub := n - 1

			for lb < ub {
				sum := nums[i] + nums[j] + nums[lb] + nums[ub]
				if sum == target {
					sol := []int{nums[i], nums[j], nums[lb], nums[ub]}
					solution = append(solution, sol)
					for lb < ub && nums[ub] == nums[ub-1] {
						ub--
					}
					ub--

					for lb < ub && nums[lb] == nums[lb+1] {
						lb++
					}
					lb++
				} else if sum > target {
					for lb < ub && nums[ub] == nums[ub-1] {
						ub--
					}
					ub--
				} else {
					for lb < ub && nums[lb] == nums[lb+1] {
						lb++
					}
					lb++
				}
			}

		}
	}

	return solution
}
```

Solution in mind
================

-	The solution builds upon 3 sum.

-	For i in the range of 0 to n-3, we calculate 3 sum between i+1 and n. We check if that sum and ith element is equal target, if yes we have a solution.
