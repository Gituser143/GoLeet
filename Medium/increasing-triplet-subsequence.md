Code
====

```go
func increasingTriplet(nums []int) bool {
	low := int(math.Exp2(31))
	high := int(math.Exp2(31))

	for _, num := range nums {
		if num <= low {
			low = num
		} else if num <= high {
			high = num
		} else {
			return true
		}
	}

	return false
}
```

Solution in mind
================

-	Make use of two variables to keep track of first and second number in triplet.

-	Traverse through the array keep track of `low` and `high`. On encountering a new minimum, update `low`. On encountering a minimum number lesser `high` but higher than `low`, update `high`.

-	This way, we encounter a third number given `low` and `high` have been set and a triplet is formed.
