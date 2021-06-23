Code
====

```go
func backspaceCompare(s string, t string) bool {
	sol := []string{}
	strs := []string{s, t}

	for _, str := range strs {
		newStr := ""
		for _, c := range str {
			if string(c) == "#" {
				if len(newStr) > 0 {
					newStr = newStr[:len(newStr)-1]
				}
			} else {
				newStr = newStr + string(c)
			}
		}

		sol = append(sol, newStr)
	}

	return sol[0] == sol[1]
}
```

Solution in mind
================

-	We iterate through characters of a string and maintain a new string (`newStr`) which has the updated string post backspaces.

-	If a `#` is encountered as the character, we simply check if `newStr` has any characters in it and remove the last character. If empty, we simply continue.

-	In the case where the encountered character is not `#`, we simply append it as a suffix to `newStr`

-	After both strings have been parsed, we compare the new strings for equality.
