Code
====

```go
func findMedianSortedArrays(left, right []int) float64 {
	n := len(left) + len(right)
	result := make([]int, n)

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	if n%2 == 0 {
		return float64(result[n/2]+result[(n/2)-1]) / 2
	} else {
		return float64(result[n/2])
	}

}
```

Solution in mind
================

-	Merge both arrays to get larger sorted array.
-	Return (nth + (n+1)th)/2
