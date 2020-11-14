Code
====

```go
func letterCombinations(digits string) []string {
	digitMap := make(map[string][]string, 0)

	digitMap["2"] = []string{"a", "b", "c"}
	digitMap["3"] = []string{"d", "e", "f"}
	digitMap["4"] = []string{"g", "h", "i"}
	digitMap["5"] = []string{"j", "k", "l"}
	digitMap["6"] = []string{"m", "n", "o"}
	digitMap["7"] = []string{"p", "q", "r", "s"}
	digitMap["8"] = []string{"t", "u", "v"}
	digitMap["9"] = []string{"w", "x", "y", "z"}

	solution := []string{}

	for _, char := range digits {
		digit := string(char)

		if len(solution) == 0 {
			solution = append(solution, digitMap[digit]...)
			continue
		}

		nextChars := digitMap[digit]

		newChars := []string{}
		for _, c := range nextChars {
			for i := range solution {
				newChars = append(newChars, solution[i]+c)
			}
		}
		solution = newChars
	}

	return solution

}
```

Solution in mind
================

-	We iterate through digits, and for each digit we create possible letter combinations with already parsed digits and append to solution.
