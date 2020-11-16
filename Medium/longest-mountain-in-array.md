Code
====

```go
func longestMountain(A []int) int {

	if len(A) < 2 {
		return 0
	}

	prev := A[0]
	sol := 0
	for i, _ := range A[1:] {
		cur := 1
		lastIndex := i
		for A[i] > prev && i < len(A)-1 {
			prev = A[i]
			i++
			cur++
		}

		if lastIndex == i {
			prev = A[i]
			continue
		}

		lastIndex = i

		for i < len(A) && A[i] < prev {
			prev = A[i]
			cur++
			i++
		}

		if lastIndex == i {
			prev = A[i]
			continue
		}

		if cur > sol && cur >= 3 {
			sol = cur
		}
	}

	return sol
}
```

Solution in mind
================

-	We iterate through the given array, if the given element is larger than the previous, we continue traversing till we reach the peak of the mountain.

-	If the traversal to the peak started and ended from the same index, we skip and try from the next element. If not, we then continue to traverse as long as the previous value is less than the current.

-	If this traversal started and ended at the same index, we skip again. If not, we check if the travelled distance is the largest travelled, if yes we update the maximum.
