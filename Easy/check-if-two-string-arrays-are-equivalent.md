Code
====

```go
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i1 := 0
	i2 := 0
	j1 := 0
	j2 := 0
	n1 := len(word1)
	n2 := len(word2)
	w1 := len(word1[0])
	w2 := len(word2[0])

	for i1 < n1 && i2 < n2 {
		if word1[i1][j1] != word2[i2][j2] {
			return false
		} else {
			j1 += 1
			j2 += 1

			if j1 == w1 {
				j1 = 0
				i1 += 1
			}

			if j2 == w2 {
				j2 = 0
				i2 += 1
			}

			if i1 == n1 || i2 == n2 {
				break
			}

			w1 = len(word1[i1])
			w2 = len(word2[i2])
		}
	}

	if i1 == n1 && i2 == n2 {
		return true
	}

	return false
}
```

Solution in mind
================

-	We make use of 4 indices, `i1, i2, j1, j2`, where `i` corresponds to the `ith` word of given subscript and the respective `j` are indices of the letter in that word.

-	We initialise all the above variables to 0. We begin a while loop where we continue to iterate through the words till we reach the last in either `word1` or `word2`.

-	In the loop, we check if the letter `word1[i1][j1] == word2[i2][j2]`. If not, we can immediately conclude that the strings do not match and we can return `false`.

-	If they do match, we have to check the next letter. This letter can either be part of the same word or the first letter in the next word.

-	We thus first increment the `j` indices and check if it has reached the end of the word (last character). If yes, we reset `j` to 0 and increment the respective `i` value so as to move to the next word in the list.

-	If either `i` exceeds the last index in corresponding word, we break from the while loop. Upon exiting the loop, we can conclude if both words match if `i1 == n1 && i2 == n2`. This condition holds if and only if, all words have been matched from both lists.
