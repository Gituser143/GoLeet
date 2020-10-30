Code
====

```go
func myAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}

	num := 0
	isNegative := 1

	for string(s[0]) == " " {
		if len(s) == 1 {
			return 0
		}
		s = s[1:]
	}

	if string(s[0]) == "-" {
		isNegative = -1
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	for _, letter := range s {
		if letter < 48 || letter > 57 {
			break
		}
		num = (num * 10) + int(letter) - 48

		if num > 2147483647 {
			if isNegative < 0 {
				return -2147483648
			} else {
				return 2147483647
			}
		}
	}

	num = num * isNegative

	return num
}
```

Solution in mind
================

-	Iterate from left to right stripping away whitespace.

-	If first letter is +/- use as sign.

-	Iterate through string and if increment solution as `sol = (sol * 10) + letter`. If the number exceeds the bounds of int32, break.

-	If a non numeric character is encountered, break.
