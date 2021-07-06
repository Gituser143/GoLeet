Code
====

```go
func minSetSize(arr []int) int {
	countMap := make(map[int]int)
	nums := []int{}
	N := len(arr)

	for _, num := range arr {
		if _, ok := countMap[num]; ok {
			countMap[num] += 1
		} else {
			countMap[num] = 1
			nums = append(nums, num)
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return countMap[nums[i]] > countMap[nums[j]]
	})

	count := 0
	for i, num := range nums {
		count += countMap[num]
		if N-count <= N/2 {
			return i + 1
		}
	}

	return N / 2
}
```

Solution in mind
================

-	Store the frequency of numbers in `countMap` and set of numbers in `nums`.

-	Sort `nums` according to frequency, in descending order.

-	Iterate over `nums` and see if removing the number, reduces the size of the array to half or lower, if yes, return the `index + 1`, as this is the number of numbers which had to be removed for `N` to be halved.
