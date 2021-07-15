Code
====

```go
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	count := 0

	for i := 0; i < n-2; i++ {
		k := i + 2
		for j := i + 1; j < n-1 && nums[i] != 0; j++ {
			for k < n && nums[i]+nums[j] > nums[k] {
				k++
			}
			count += k - j - 1
		}
	}

	return count
}
```

Solution adopted from [here](https://leetcode.com/problems/valid-triangle-number/solution/)
