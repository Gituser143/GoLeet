Code
====

```go
func findKthPositive(arr []int, k int) int {
	missing := []int{}
	l := 0
	prev := arr[0]

	for i := 1; i < prev && l <= k; i++ {
		missing = append(missing, i)
		l += 1
	}

	if l >= k {
		return missing[k-1]
	}

	for _, num := range arr[1:] {
		for i := prev + 1; i < num && l <= k; i++ {
			missing = append(missing, i)
			l += 1
		}

		if l >= k {
			return missing[k-1]
		}

		prev = num
	}

	if k-l > 0 {
		return arr[len(arr)-1] + k - l
	}

	return missing[k-1]
}
```

Solution in mind
================

-	We keep an integer slice (`missing`) to keep track of missing elements. We maintain two integers, `l` to keep track of length of `missing` and `prev`, to track last seen element from input array.

-	First get all missing elements from 1 up to first element in the array. This is the range `[1,arr[0])`. These elements are stored in `missing` (This is skipped if `arr[0] == 1`).

-	We initialise `prev = arr[0]` following which we iterate through `arr`, skipping arr[0].

-	We check if the number `num` and `prev` are adjacent, if not there are values between them that are missing. We populate `missing` by adding values from the range `(prev, num)`.

-	On each append, we increment `l`. When `l` reaches the value `k`, we have `k` missing elements and can return the answer (`missing[k-1]`).

-	If we iterate through the array and have not populated `k` missing elements. The missing element will occur after the last element in the input array. Thus we return `lastElement + k - l`. We subtract `l` from this because we might have already seen a few missing elements.
