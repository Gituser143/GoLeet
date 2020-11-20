Code
====

```go
func search(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
```

Solution in mind
================

-	Naive linear search
