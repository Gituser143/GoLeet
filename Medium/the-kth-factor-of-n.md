Code
====

```go
func kthFactor(n int, k int) int {
	j := 0
	for i := 1; j <= k; i++ {
		if n%i == 0 {
			j++
		}

		if j == k {
			return i
		} else if i > n {
			return -1
		}
	}

	return -1
}
```

Solution in mind
================

-	To find the kth factor of `n`, we use two variables i and j.

-	We loop using `i` and use `j` to keep count of the nmber of factors encountered. We encounter a factor when `n % i == 0`, i being the factor of n.

-	On encountering a factor, we increment the value of `j` indicating we have found a new factor.

-	When the value of `j` is the same as `k`, we now have the kth factor of n, which is the value in `i`.

-	The kth factor doesn't exist when `i` exceeds `n`.
