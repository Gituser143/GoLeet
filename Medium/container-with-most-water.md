Code
====

```go
func maxArea(height []int) int {
	if len(height) == 2 {
		if height[0] < height[1] {
			return height[0]
		} else {
			return height[1]
		}
	}

	i := 0
	j := len(height) - 1

	maxCapacity := 0
	for i != j {
		h := 0
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}

		capacity := h * (j - i + 1)

		if capacity > maxCapacity {
			maxCapacity = capacity
		}
	}
	return maxCapacity
}
```

Solution in mind
================

-	Keep two pointers starting from beginning and end.

-	Calculate area covered by these pointers and if greater than max, update.

-	Move the pointer with the lower value in height towards other pointer.

-	Repeat area calculation. Loop until two pointers meet.

-	Return max area
