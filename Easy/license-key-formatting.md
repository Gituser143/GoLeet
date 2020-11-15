Code
====

```go
func licenseKeyFormatting(S string, K int) string {
	s := strings.ReplaceAll(S, "-", "")
	s = strings.ToUpper(s)

	n := len(s)

	sol := ""
	k := K

	fmt.Println(s)

	if n%k == 0 {
		for i := k; i <= n; i += k {
			sol += s[:k]
			sol += "-"
			s = s[k:]
		}
	} else {
		sol += s[:n%k] + "-"
		x := (n % k)
		s = s[x:]
		fmt.Println(s, x)

		n -= x
		for i := k; i <= n; i += k {
			sol += s[:k]
			sol += "-"
			s = s[k:]
		}
	}

	if len(sol) == 0 {
		return sol
	}

	return sol[:len(sol)-1]
}
```

Solution in mind
================

-	We first remove all occurrences of "-" in the string and then convert it to upper case.

-	We then check if the first group would be able to fit k strings. We can figure this out if n % k is equal to 0. If yes, we iteratively join the string in hunks of k groups separated by a "-".

-	If n % k is not 0, we make the first (n % k) characters into one group and split the rest on chunks of k.
