Code
====

```go
func maximalNetworkRank(n int, roads [][]int) int {
	rankMap := make(map[int]map[int]bool, 0)
	for _, road := range roads {
		city1, city2 := road[0], road[1]

		if _, ok := rankMap[city1]; ok {
			rankMap[city1][city2] = true
		} else {
			rankMap[city1] = map[int]bool{city2: true}
		}

		if _, ok := rankMap[city2]; ok {
			rankMap[city2][city1] = true
		} else {
			rankMap[city2] = map[int]bool{city1: true}
		}
	}

	solution := 0

	for i := 0; i < n; i++ {
		if _, ok := rankMap[i]; !ok {
			continue
		}
		for j := i + 1; j < n; j++ {
			if city2, ok := rankMap[j]; !ok {
				continue
			} else {
				rank := len(rankMap[i]) + len(rankMap[j])
				if _, ok := city2[i]; ok {
					rank = rank - 1
				}

				if rank > solution {
					solution = rank
				}
			}
		}
	}

	return solution
}
```

Solution in mind
================

-	Implement a map `rankMap` to store `city: set of adjacent cities`. Set is another map of type `city: bool`.

-	Iterate over roads and update values of `rankMap` accordingly.

-	Once rankMap is ready, iterate `i` from 0 to `n-1` and `j` from `i+1` to `n`. If city i does not exist in map, continue. If it does, check the same for city j.

-	Set rank as len(setCity1) + len(setCity2).

-	If city1 exists in setCity2, they have a common road, so reduce rank by 1.

-	Verify if this rank is greater than previous and update solution.
