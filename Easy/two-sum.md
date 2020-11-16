Code
====

```go
func twoSum(nums []int, target int) []int {
	compliment := make(map[int]int, 0)

	for i, val := range nums {
		com := target - val

		if j, ok := compliment[com]; ok {
			return []int{i, j}
		}

		compliment[val] = i
	}

	return []int{0, 0}
}
```

Solution in mind
================

-	We make use of map to store index: number.

-	We iterate through the array and check to see if we have it's compliment (target - current). If yes, we return the indices of these two number.

-	If the compliment does not exist, we just add the current number to the map.
