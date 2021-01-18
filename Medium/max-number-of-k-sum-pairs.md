Code
====

```go
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxOperations(nums []int, k int) int {
	numMap := make(map[int]int, 0)

	for _, num := range nums {
		if num < k {
			if _, ok := numMap[num]; ok {
				numMap[num] += 1
			} else {
				numMap[num] = 1
			}
		}
	}

	ops := 0

	for num, count := range numMap {
		compliment := k - num

		if num == compliment {
			ops += count / 2
			numMap[num] = count % 2
		} else {
			if cCount, ok := numMap[compliment]; ok {
				x := min(cCount, count)
				numMap[compliment] -= x
				numMap[num] -= x
				ops += x
			}
		}

		if numMap[compliment] == 0 {
			delete(numMap, compliment)
		}

		if numMap[num] == 0 {
			delete(numMap, num)
		}

	}

	return ops
}
```

Solution in mind
================

-	We start by populating a map (`numMap`) of format `num: count`.

-	We only track numbers with values lesser than `k` as values greater than k are irrelevant (cannot form a pair with `sum = k`).

-	We iterate through the numbers of `numMap` as `num` with corresponding count `count`. For each number, we check if it's `compliment`, `k - num` exists.

-	If for a number, it's `compliment` exists, we find the lesser of their counts and subtract both their counts with this value. Similarly we add this value to `ops` which stores the number of operations performed.

-	We pick the lesser of the two counts as that is the number of pairs that can be formed (using `num` and `compliment`).

-	A special case arises when `num` == `compliment`. Here, if the `count` is even, all numbers can be paired up and removed in `count / 2` operations. If `count` is odd, `count - 1` numbers can be removed in `count / 2` operations and the single number will be left out.

-	At the end of iterating through the map, we return `ops`.
