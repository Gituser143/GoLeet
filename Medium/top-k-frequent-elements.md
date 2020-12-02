Code
====

```go
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]map[int]bool, 0)
	set := make(map[int]int, 0)
	maxCount := 0

	for _, num := range nums {

		if count, ok := set[num]; ok {
			set[num] += 1
			delete(countMap[count], num)

			if _, ok := countMap[count+1]; !ok {
				numSet := make(map[int]bool, 0)
				countMap[count+1] = numSet
			}

			countMap[count+1][num] = true

			if count+1 > maxCount {
				maxCount = count + 1
			}

		} else {
			set[num] = 1

			if _, ok := countMap[1]; ok {
				countMap[1][num] = true
			} else {
				numSet := make(map[int]bool, 0)
				countMap[1] = numSet
				countMap[1][num] = true
			}

			if 1 > maxCount {
				maxCount = 1
			}

		}

	}

	sol := []int{}

	for i := 0; i < k; i++ {
		if numbers, ok := countMap[maxCount]; ok {
			for num, _ := range numbers {
				sol = append(sol, num)
				i++
			}
			i--
		}

		maxCount--
	}

	return sol
}
```

Solution in mind
================

-	We first create 2 data structures:

	-	A `countMap` which holds `count: set_of_nums()`, this holds a `count`, and all the numbers which occur `count` number of times.

	-	A `set` which is a set to check if a number already exists in `countMap`. This keeps track of the number and the number of times it was seen.

-	We keep a variable `maxCount` to keep track of max occurred count.

-	We iterate through `nums`, for each number `num` in `nums`, we check if that number has occurred previously, by checking if it exists in the set. If it does, we have to remove it from the previous `count` list in from `countMap`. Following which we place it at the `count + 1` list.

-	If the number is encountered for the first time, we add it to `count[1]` and initialise it in `set` too.

-	As and when a number is added to `countMap`, we check if that `count` is the maximal encountered count. If yes, we update maxCount.

-	Once we have populated `countMap`, we now have the `maxCount` too. We then iterate from `0 to k` using a variable `i`, getting numbers of count `maxCount`. We decrement maxCount each iteration.

-	We iterate through above list, incrementing `i` for each encountered number in list. We decrement `i` 1 time after the list iteration. We do this to ensure unnecessary increments when only 1 element was added to Solution list or when no element was added.
