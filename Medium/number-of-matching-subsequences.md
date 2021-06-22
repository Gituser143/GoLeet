Code
====

```go
func numMatchingSubseq(s string, words []string) int {
	strMap := make(map[rune][]int)

	// Initialise strMap
	for i, c := range s {
		if val, ok := strMap[c]; ok {
			strMap[c] = append(val, i)
		} else {
			strMap[c] = []int{i}
		}
	}

	solCount := 0

	for _, word := range words {
		i := 0
		var c rune
		prevIdx := -1

		for i, c = range word {
			// Check if letter exists in map
			if idxs, ok := strMap[c]; ok {
				idxFound := false

				// Find index of letter > than index of previously found letter
				for _, idx := range idxs {
					if idx > prevIdx {
						prevIdx = idx
						idxFound = true
						break
					}
				}

				// If no valid index is found, break
				if !idxFound {
					i = 0
					break
				}
			} else {
				i = 0
				break
			}

		}

		// Word is a subsequence if all letters were successfully found in strMap
		if i == len(word)-1 {
			solCount += 1
		}
	}

	return solCount
}
```

Solution in mind
================

-	Create a map (`strMap`) which stores letters of given string (`s`) as key and all indexes of letter in string as value (stored as `[]int`).

-	Iterate over each word, for each word, validate if it is a subsequence.

-	To validate if word is subsequence, iterate over letters of word, if letter does not exist in `strMap`, word is not subsequence.

-	If word exists, find the least index of it's occurrence such that it is greater than the index of the previous word. If no such index is found, the relative order of letters can't be maintained and the word is not a subsequence.

-	Validate word as a subsequence only if all letters were successfully found in the `strMap` (`i == len(word)-1`).

-	Return count of validated words (`solCount`).

