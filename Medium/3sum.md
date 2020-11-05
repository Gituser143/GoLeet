Code
====

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	solution := [][]int{}

	for i := 0; i < n-2; i++ {
		for i != 0 && nums[i] == nums[i-1] && i < n-2 {
			i++
		}

		if nums[i] > 0 {
			break
		}

		lowerB := i + 1
		upperB := n - 1

		for lowerB < upperB {
			sum := nums[i] + nums[lowerB] + nums[upperB]
			if sum == 0 {
				solution = append(solution, []int{nums[i], nums[lowerB], nums[upperB]})

				for lowerB < upperB && nums[upperB] == nums[upperB-1] {
					upperB--
				}
				upperB--

				for lowerB < upperB && nums[lowerB] == nums[lowerB+1] {
					lowerB++
				}
				lowerB++

			} else if sum < 0 {
				for lowerB < upperB && nums[lowerB] == nums[lowerB+1] {
					lowerB++
				}
				lowerB++
			} else {
				for lowerB < upperB && nums[upperB] == nums[upperB-1] {
					upperB--
				}
				upperB--
			}
		}
	}

	return solution
}
```

Solution in mind
================

-	First we sort the given array in ascending order.

-	To get 3 numbers who's sum is 0, we turn the problem into finding 1 number and a pair who's sum is 0.

-	We iterate through the rest of the array and try to find a pair of numbers such that the sum of all 3 is 0.

-	To find this pair we keep a lowerBound index as i+1 and upperBound index as n-1.

-	If we find one, we add the three numbers to a solution list.

-	If the sum is lesser than 0, we increment the lower bound. We continue to increment the lower bound till two adjacent values are not equal. This helps avoiding duplicates.

-	Similarly if the sum is greater than 0, we continually decrement the upperBound till we get a number different from current upperBound.

-	We repeat this for every i in the list and break if i is a positive number, as all numbers after i will only give sums greater than 0.
