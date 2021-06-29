Code
====

```go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func kSimilarity(s1 string, s2 string) int {
	A := []byte(s1)
	B := []byte(s2)

	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			continue
		}

		matches := []int{}

		for j := i + 1; j < len(A); j++ {
			if A[j] == B[i] && A[j] != B[j] {
				matches = append(matches, j)
				if A[i] == B[j] {
					A[i], A[j] = A[j], A[i]
					return 1 + kSimilarity(string(A[i+1:]), string(B[i+1:]))
				}
			}
		}

		best := len(A) - 1

		for _, j := range matches {
			A[i], A[j] = A[j], A[i]
			best = min(best, 1+kSimilarity(string(A[i+1:]), string(B[i+1:])))
			A[i], A[j] = A[j], A[i]
		}
		return best
	}

	return 0
}
```

Solution in mind
================

Solution borrowed from [here](https://leetcode.com/problems/k-similar-strings/discuss/140299/C%2B%2B-6ms-Solution).

\-
