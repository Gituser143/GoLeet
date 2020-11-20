Code
====

```go
func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
```

Solution in mind
================

-	Naive linear search
