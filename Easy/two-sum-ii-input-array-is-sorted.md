Code
====

```go
func twoSum(numbers []int, target int) []int {

	nums := numbers

	lb := 0
	ub := len(nums) - 1

	for ub > lb {
		sum := nums[ub] + nums[lb]

		if sum == target {
			return []int{lb + 1, ub + 1}
		} else if sum < target {
			for lb < ub && nums[lb+1] == nums[lb] {
				lb++
			}
			lb++
		} else {
			for lb < ub && nums[ub] == nums[ub-1] {
				ub--
			}
			ub--
		}
	}

	return []int{0, 0}
}
```

Solution in mind
================

-	Since the array i sorted, we make use of two pointers, upper bound (ub) and lower bound (lb).

-	We initialise ub, lb to `len(nums)-1` and 0 respectively. We check if the sum of elements at these indices are equal to target. If yes, we return indices.

-	If the sum is lesser, we move the lower bound pointer higher up. If greater, we bring the upper bound pointer down.
