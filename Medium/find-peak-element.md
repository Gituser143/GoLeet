Code
====

Naive Solution
--------------

```go
func findPeakElement(nums []int) int {
	n := len(nums) - 1

	if n == 0 {
		return 0
	} else if n == 1 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	for i, num := range nums {
		if i == 0 || i == n {
			continue
		}
		if num > nums[i-1] && num > nums[i+1] {
			return i
		}
	}

	if nums[n] > nums[n-1] {
		return n
	}

	return 0
}
```

Using Binary Search
-------------------

```go
func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
```
