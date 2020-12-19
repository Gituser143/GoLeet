Code
====

```go
func fourSumCount(A []int, B []int, C []int, D []int) int {
	count := 0
	sums := make(map[int]int, 0)

	for _, a := range A {
		for _, b := range B {
			if _, ok := sums[a+b]; ok {
				sums[a+b] += 1
			} else {
				sums[a+b] = 1
			}
		}
	}

	for _, c := range C {
		for _, d := range D {
			if val, ok := sums[-1*(c+d)]; ok {
				count += val
			}
		}
	}

	return count
}
```

Solution in mind
================

-	Create a map `sums` which stores count of sum occurrences of elements in first two arrays.

-	Populate it by computing sums from `A` and `B`.

-	Following which, iterate through `C` and `D` and check if the sum of a pair of elements complements a sum already existing in `sums`, if it does, it satisfies the `a + b + c + d = 0`.

-	Increment count on encountering sum = 0.
