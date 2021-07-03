Code
====

```go
func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	} else if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	sol := []int{}
	idx := -1
	var i, num int

	// Find number closest to x and its index
	for i, num = range arr {
		if i == len(arr)-1 {
			break
		}
		if num == x || (arr[i] < x && arr[i+1] > x) {
			idx = i
			if num != x {
				if x-arr[i] > arr[i+1]-x {
					num = arr[i+1]
					idx = i + 1
				}
			}
			break
		}
	}

	sol = append(sol, num)

	count := 1
	front := 1
	back := 1

	for count < k {
		// If both front and back are within range
		if idx-back >= 0 && idx+front < len(arr) {
			// If both values are same, use number at lower index
			if arr[idx+front] == arr[idx-back] {
				sol = append([]int{arr[idx-back]}, sol...)
				back += 1
			} else {
				// Use closer number
				if x-arr[idx-back] <= arr[idx+front]-x {
					sol = append([]int{arr[idx-back]}, sol...)
					back += 1
				} else {
					sol = append(sol, arr[idx+front])
					front += 1
				}
			}
			count += 1
		} else {
			// Add all numbers from either left or right side of the window
			if idx+front >= len(arr) {
				sol = append([]int{arr[idx-back]}, sol...)
				count += 1
				back += 1
			} else {
				sol = append(sol, arr[idx+front])
				front += 1
				count += 1
			}
		}
	}

	return sol
}
```

Solution in mind
================

-	Find a number closest to `x` and locate it's index as `idx`.

-	Find all number on the left and right of it according to given conditions of closeness and get `k` of numbers.

-	The `k` numbers are found by maintaining a window of adjacent numbers following the closeness conditions.

-	The window is returned when `k` numbers have been put into it.
