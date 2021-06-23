Code
====

```go
func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func maxAscendingSum(nums []int) int {
	window := []int{}
	windowSum := 0
	prevNum := 0

	for _, num := range nums {
		if num > prevNum {
			window = append(window, num)
			prevNum = num
		} else {
			s := sum(window)
			if s > windowSum {
				windowSum = s
			}
			prevNum = num
			window = []int{num}
		}
	}

	s := sum(window)
	if s > windowSum {
		windowSum = s
	}

	return windowSum
}
```

Solution in mind
================

-	We iterate through elements of the `nums` array, while keeping track of `window` and `prevNum`. Here, `window` holds the subsequence of the array that is ascending.

-	As and when we encounter a number, we check if it is greater than the previous encountered number. If it is, the `window` grows in size (we append to `window`).

-	If it is not, we calculate the sum of the existing window and check if it is greater than any previously encountered sum (`windowSum`) and update the sum accordingly.

-	After updating the window sum, we reset the window to have only the latest encountered number in it, as the window is no longer increasing.

-	After completion of iterations through elements, we return windowSum
