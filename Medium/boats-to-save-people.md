Code
====

```go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numRescueBoats(people []int, limit int) int {
	weights := make(map[int]int, 0)
	weightSet := []int{}

	for _, weight := range people {
		if _, ok := weights[weight]; !ok {
			weights[weight] = 1
			weightSet = append(weightSet, weight)
		} else {
			weights[weight] += 1
		}
	}

	numBoats := 0

	for _, weight := range weightSet {
		if weight >= limit {
			numBoats += weights[weight]
			delete(weights, weight)
			continue
		}

		complimentWeight := limit - weight

		for i := complimentWeight; i > 0; i-- {
			if i == weight {
				numBoats += weights[weight] / 2
				weights[weight] = weights[weight] % 2

				if weights[weight] == 0 {
					break
				} else {
					continue
				}
			}

			if _, ok := weights[i]; ok {
				minCount := min(weights[i], weights[weight])
				weights[i] -= minCount
				weights[weight] -= minCount
				numBoats += minCount
			}

			if weights[weight] == 0 {
				break
			}
		}

		if weights[weight] != 0 {
			numBoats += weights[weight]
			delete(weights, weight)
		} else {
			delete(weights, weight)
		}
	}

	return numBoats
}
```

Solution in mind
================

-	We first create a map called `weights`, which holds data in the format `weight: count`, where weight is the weight of a person and count is the count of people with that weight.

-	Following this, we iterate through unique weights.

-	If the `weight` is above the given `limit`, we will need 1 boat for each person of said weight.

-	If the `weight` is lesser than `limit`, we must find a `complimentWeight` such that the sum of these two weights are lesser than `limit.`

-	Upon finding a `complimentWeight`, we can send `x` pairs of people on `x` boats. The number `x` is determined by the minimum number of people with weight = `weight` or `complimentWeight`.

-	A key point to note is that `complimentWeight + weight` need not have to be equal to `limit`. The sum being lesser works too. Thus, we check for compliment weights in the range `[1, complimentWeight]`.

-	Finally, if a `weight` does not have a compliment, or has a higher count than any possible compliment, the remainder (or entire count if no compliment found), must be sent through single boats.
