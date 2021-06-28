Code
====

```go
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func candy(ratings []int) int {
	n := len(ratings)

	// Initialise left-to-right and right-to-left arrays
	l2r := make([]int, n)
	r2l := make([]int, n)
	for i := 0; i < n; i++ {
		l2r[i] = 1
		r2l[i] = 1
	}

	// Populate arrays
	for i := 0; i < n-1; i++ {
		j := n - 1 - i

		// Populate left-to-right array
		if ratings[i+1] > ratings[i] {
			l2r[i+1] = l2r[i] + 1
		}

		// Populate right-to-left array
		if ratings[j-1] > ratings[j] {
			r2l[j-1] = r2l[j] + 1
		}
	}

	// Find required num of candies
	sum := 0
	for i := 0; i < n; i++ {
		sum += max(l2r[i], r2l[i])
	}

	return sum
}
```

Solution in mind
================

Adopted from Approach 2 from [leetcode](https://leetcode.com/problems/candy/solution/).

-	We maintain 2 arrays (`l2r`, `r2l`) to hold candies.

-	The `l2r` array holds candies only with respect to left-to-right comparisons. Thus, if the rating of a child on the right is higher than a child on the left, the right child gets left child's candies + 1.

-	Similarly, `r2l` array holds candies with respect to right-to-left comparisons.

-	We can then find the total number of candies required, by summing the max element at every index from both the arrays.
