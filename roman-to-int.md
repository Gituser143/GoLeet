Code
====

```go
func romanToInt(s string) int {
	// Init map to hold numbers
	num := make(map[string]int, 0)
	num["I"] = 1
	num["V"] = 5
	num["X"] = 10
	num["L"] = 50
	num["C"] = 100
	num["D"] = 500
	num["M"] = 1000

	sol := 0

	// Iterate through letters
	for i, letter := range s {
		// Increment solution by given number
		sol += num[string(letter)]

		// If not the first number
		if i > 0 {
			char := string(s[i-1])
			// If previous number used to reduce current number
			// Like IV for 4 and not 1+5
			if char == "I" || char == "X" || char == "C" {
				current := string(letter)
				// Subtract the previous number added
				switch char {
				case "I":
					if current == "V" || current == "X" {
						sol -= 2 * num[char]
					}
				case "X":
					if current == "L" || current == "C" {
						sol -= 2 * num[char]
					}
				case "C":
					if current == "D" || current == "M" {
						sol -= 2 * num[char]
					}
				}
			}
		}
	}

	return sol
}
```

Solution in mind
================

-	Iterate over letters.
-	Add each letter.
-	if previous letter is used to reduce current letter, subtract previous letter twice.
