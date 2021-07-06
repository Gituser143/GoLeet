Code
====

```go
// getSums returns all possible subset sums for a given slice of nums
func getSums(nums []int) []int {
	n := len(nums)
	total := 1 << n
	sums := []int{}

	for i := 0; i < total; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			mask := 1 << j
			if i&mask != 0 {
				sum += nums[j]
			}
		}
		sums = append(sums, sum)
	}

	return sums
}

// binarySearch finds index of element closest to x in nums
func binarySearch(nums []int, x int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= x })
}

// min calculates the lesser of a and abs(b)
func min(a, b int) int {
	// make b positive if negative
	if b < 0 {
		b *= -1
	}

	// return smaller value
	if a < b {
		return a
	}

	return b
}

func minAbsDifference(nums []int, goal int) int {
	diff := 10000000000

	// split the given nums into two halves and compute their sums
	mid := len(nums) / 2
	l1 := getSums(nums[:mid])
	l2 := getSums(nums[mid:])

	// Sort the second half in ascending order
	sort.Ints(l2)

	// for each number in l1, find it's compliment such that
	// l1-compliment is closest to goal
	for _, num := range l1 {
		// If the number is goal or the previously computed
		// difference is 0, we can return
		if num == goal || diff == 0 {
			return 0
		}

		// Get compliment and it's index in l2
		compliment := goal - num
		complimentIdx := binarySearch(l2, compliment)

		// Since complimentIdx is the index of number larger or equal
		// to the given compliment, we check for both the index and the index
		// before it, as the number at the index before it might be closer
		// than the number at the index
		if complimentIdx < len(l2) {
			diff = min(diff, compliment-l2[complimentIdx])
		}

		if complimentIdx > 0 {
			diff = min(diff, compliment-l2[complimentIdx-1])
		}

	}

	return diff
}
```

Solution in mind
================

-	We split the given `nums` array into two equal (almost) halves and generate all possible subset sums for the halves.

-	We sort one of the two subset sums.

-	We then iterate over the unsorted subset and attempt to find a compliment for the number. The compliment is `goal - num` as `num + compliment = goal`.

-	We find this compliment using binary search. Following which, we compute the difference and update `diff`.
