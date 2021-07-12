Code
====

```go
func isIsomorphic(s string, t string) bool {
	charMap := make(map[byte]byte)
	mappedSet := make(map[byte]bool)

	n := len(s)

	for i := 0; i < n; i++ {
		val, ok := charMap[s[i]]
		if ok && val == t[i] {
			continue
		} else if ok {
			return false
		} else {
			if _, ok := mappedSet[t[i]]; ok {
				return false
			} else {
				mappedSet[t[i]] = true
				charMap[s[i]] = t[i]
			}
		}
	}

	return true
}
```

Solution in mind
================

-	We keep track of a map called `charMap` which holds the records of mappings between letters from `s` to `t`.

-	We also keep a set called `mappedSet` which holds the boolean value of whether a character has already been mapped to.

-	We iterate from `i = 0` to `i = len(s)`. We check if the characters at `s[i]` already exist in the `charMap`, if yes and if the mapping is equal to the character at `t[i]`, we can continue our checks.

-	But if a mapping exists and the mapped value is not equal to the character at `t[i]`, we can immediately return `false` as a previous mapping prevents the current `s[i]` to be mapped to a new character.

-	Finally, if no mapping exists, we attempt to create a mapping. But in order to do so, we must ensure the `t[i]` character has not already been mapped to. We do this by checking if it exists in the `mappedSet`. If yes, we return `false`, else we can create the mapping.

-	If we finish the iterations, we can conclude that the strings are isomorphic.
