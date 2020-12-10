Code
====

```go
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	n := len(arr)

	i := 0
	for i = 0; i < n-1; i++ {
		if arr[i] >= arr[i+1] {
			break
		}
	}

	if i == 0 {
		return false
	}

	j := i
	for j = i + 1; j < n; j++ {
		if arr[j] >= arr[j-1] {
			return false
		}
	}

	if j == i+1 {
		return false
	}

	return true
}
```

Solution in mind
================

-	Solution is a simple one pass traversal of the array.

-	As we climb p the array we ensure that `arr[i+1] >= arr[i]`. When this is violated, we reach a peak.

-	We must ensure the peak is not at the very beginning (`i == 0`), if it was, we never climbed up the mountain.

-	After reaching peak, we begin to climb down, while ensuring the condition `arr[i] < arr[i-1]`, if not, we either have a rise or flat zone in the mountain.

-	After descending from the peak, we ensure we have actually moved downwards and not stayed at the peak (`j == i+1`). This acts as a check to ensure the peak is not at the very end of the array too.
