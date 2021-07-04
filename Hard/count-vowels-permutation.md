Code
====

```go
func countVowelPermutation(n int) int {
	nextMap := map[string][]string{
		"":  {"a", "e", "i", "o", "u"},
		"a": {"e"},
		"e": {"a", "i"},
		"i": {"a", "e", "o", "u"},
		"o": {"i", "u"},
		"u": {"a"},
	}

	count := 0
	prevLetters := map[string]int{"": 1}
	currLetters := make(map[string]int)

	for i := 0; i < n; i++ {
		currLetters = map[string]int{
			"a": 0,
			"e": 0,
			"i": 0,
			"o": 0,
			"u": 0,
		}

		for letter, num := range prevLetters {
			for _, nextLetter := range nextMap[letter] {
				currLetters[nextLetter] = (currLetters[nextLetter] + num) % 1000000007
			}
		}

		prevLetters = currLetters
	}

	for _, num := range currLetters {
		count = (count + num) % 1000000007
	}

	return count
}
```

Solution in mind
================

-	We maintain a map (`nextMap`) of possible next vowels given a vowel.

-	We start with a map `prevLetters` and `currLetters`. These maps hold the count of a letter.

-	As we only need the last letter of a word to generate the next set of combinations, we map the count of words ending with a letter to the letter in `currLetters`.

-	`currLetters` is populated by iterating through the previous set of letters and updating vowel counts.

-	We finally return `count` as the sum of vowel counts in `currLetters`.
