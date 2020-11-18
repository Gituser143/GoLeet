Code
====

```go
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	sol := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		added := false
		for _, solInterval := range sol {
			if interval[0] >= solInterval[0] && interval[0] <= solInterval[1] {
				if interval[1] > solInterval[1] {
					solInterval[1] = interval[1]
				}
				added = true
			}
		}
		if !added {
			sol = append(sol, interval)
		}
	}
	return sol
}
```

Solution in mind
================

-	We start off by sorting the intervals slice based off the first element in each interval. This way we get ascending intervals.

-	We then add the first interval into the solution slice.

-	We iterate through the rest of the intervals and check if the first element of the interval first in between any already existing interval in the solution.

-	If it does, we update the second element in the solution interval to the maximum of second elements between the interval in solution and the one in intervals.

-	If it does not, we append this interval to the list of solution intervals.
