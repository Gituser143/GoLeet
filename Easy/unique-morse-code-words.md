Code
====

```go
func uniqueMorseRepresentations(words []string) int {
	codes := make(map[string]string, 0)
	codes["a"] = ".-"
	codes["b"] = "-..."
	codes["c"] = "-.-."
	codes["d"] = "-.."
	codes["e"] = "."
	codes["f"] = "..-."
	codes["g"] = "--."
	codes["h"] = "...."
	codes["i"] = ".."
	codes["j"] = ".---"
	codes["k"] = "-.-"
	codes["l"] = ".-.."
	codes["m"] = "--"
	codes["n"] = "-."
	codes["o"] = "---"
	codes["p"] = ".--."
	codes["q"] = "--.-"
	codes["r"] = ".-."
	codes["s"] = "..."
	codes["t"] = "-"
	codes["u"] = "..-"
	codes["v"] = "...-"
	codes["w"] = ".--"
	codes["x"] = "-..-"
	codes["y"] = "-.--"
	codes["z"] = "--.."

	seenWords := make(map[string]bool, 0)

	for _, word := range words {
		current := ""
		for _, char := range word {
			letter := string(char)
			current += codes[letter]
		}

		if _, ok := seenWords[current]; !ok {
			seenWords[current] = true
		}
	}

	return len(seenWords)

}
```

Solution in mind
================

-	Maintain a set of seen words. And a map to map letter to it's morse equivalent

-	Iterate through words and add it's morse translation into the set.

-	Since the set can have only unique vales, return the length of the set.
