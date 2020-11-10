Code
====

```go
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	n := len(nums)

	closestDif := target - (nums[0] + nums[1] + nums[2])
	if closestDif < 0 {
		closestDif *= -1
	}

	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		lb := i + 1
		ub := n - 1
		current := nums[i]

		for lb < ub {
			sum := current + nums[lb] + nums[ub]

			if sum == target {
				return target
			} else if sum < target {
				lb++
			} else {
				ub--
			}

			dif := target - sum
			if dif < 0 {
				dif *= -1
			}

			if dif < closestDif {
				closestDif = dif
				closestSum = sum
			}
		}
	}

	return closestSum
}
```

Solution in mind
================

-	Iterated through sorted array and made use of two pointers.

-	We move the two pointers from index i + 1 and the end of array till they meet.

-	We calculate 3 sums for current, upper bound and lower bound. If the difference the sum and the target is lower than previous encountered sums, we update closestSum.

-	If the sum is lesser than the target we move the lower bound pointer one position rightwards. If sum is greater, we move the upper bound pointer one position leftwards.
