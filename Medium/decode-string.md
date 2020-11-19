Code
====

```go
func decodeString(s string) string {
	i := len(s) - 1

	for {
		if i < 0 {
			return s
		}

		if unicode.IsDigit(rune(s[i])) {
			strBeg := i + 2
			num := ""
			for i >= 0 && unicode.IsDigit(rune(s[i])) {
				num = string(s[i]) + num
				i--
			}
			beg := i
			i = strBeg
			n, _ := strconv.Atoi(num)
			fmt.Println(num)
			str := ""
			for string(s[i]) != "]" {
				str += string(s[i])
				i++
			}
			end := i

			s = strings.ReplaceAll(s, s[beg+1:end+1], strings.Repeat(str, n))

			i = beg
		}
		i--
	}
}
```

Solution in mind
================

-	We start at the end of string and move index by index backwards till we encounter a number. When we encounter a number if nested, it will be at the inner most depth.

-	On encountering a number we continue to move back till we get a non digit or reach the 0th index. This is to handle multi digit numbers.

-	Once we've parsed the full number we increment the index from the beginning of string (index of first encountered number + 2) till we encounter a closing square bracket "]".

-	We now have the string to repeat, the number of times to repeat it for and the string to replace by this new repeated string. We thus proceed to replace the encoded string (format: number[str]), with the repeated string.

-	We then continue to move the index backwards from the position of the last digit of the number, aka the beginning index of the replaced string.

-	We return the string when we have fully parsed the string.
