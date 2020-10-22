Code
====

```go
func findJudge(N int, trust [][]int) int {
	truster := make([]bool, N)           // Person : [does he trust anyone]
	trusteeMap := make(map[int][]int, 0) // Person : [trusters]

	if len(trust) == 0 && N == 1 {
		return 1
	}

	for _, val := range trust {
		if _, ok := trusteeMap[val[1]]; ok {
			trusteeMap[val[1]] = append(trusteeMap[val[1]], val[0])
		} else {
			trusteeMap[val[1]] = []int{val[0]}
		}
		truster[val[0]-1] = true
	}

	for key, val := range trusteeMap {
		if len(val) == N-1 {
			if !truster[key-1] {
				return key
			}
		}
	}
	return -1
}
```

Solution in mind
================

-	Iterate through trust array and keep track of `person: [people trusting person]`.

-	Keep a second map to see if a person trusts someone.

-	Iterate through trusteeMap map to find a person everyone trusts. If this person does not trust anyone, this person is the judge.
