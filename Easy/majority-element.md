Code
====

```go
func majorityElement(nums []int) int {
	n := int(len(nums) / 2)

	count := make(map[int]int, 0)

	for _, num := range nums {
		if _, ok := count[num]; ok {
			count[num] += 1
		} else {
			count[num] = 1
		}
	}

	for num, c := range count {
		if c > n {
			return num
		}
	}
	return 0
}
```

Solution in mind
================

-	Iterate through given array and keep track of count of numbers in a map (number: count).

-	Iterate through the map and return the first number which has a number greater than floor(n/2).
