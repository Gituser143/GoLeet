Code
====

Naive Solution
--------------

```go
func subarraySum(nums []int, k int) int {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == k {
			count += 1
		}
		sum := nums[i]
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			if sum == k {
				count += 1
			}
		}
	}

	return count
}
```

Using Hashmap
-------------

```go
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for _, num := range nums {
		sum += num
		if c, ok := sumMap[sum-k]; ok {
			count += c
		}
		if _, ok := sumMap[sum]; ok {
			sumMap[sum] += 1
		} else {
			sumMap[sum] = 1
		}
	}

	return count
}
```

Solution in mind
================

Naive solution
--------------

-	Here, we iterate through each element in the `nums` array.
-	For the given number, we try to see if any subarray starting from that number, till the end gives a sum equal to `k`. If yes, we increment `count`.

Hashmap solution
----------------

-	We maintain a map (`sumMap`), which stores `sum: number of occurrences of sum`.
-	We then iterate through the array and store cumulative sums in the map.
-	If a given `cumulative sum - k` exists in the map, then we have a subarray with sum `k` where the subarray starts at the end of where the cumulative sum was computed and ends at the current number we are at.
-	We thus increment the count by the number of ways `cumulative sum-k` has occurred.
