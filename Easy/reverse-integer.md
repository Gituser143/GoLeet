Code
====

```go
func reverse(x int) int {
	sol := 0
	for x != 0 {
		num := x % 10
		sol = sol * 10
		sol += num
		x = x / 10
	}

	if float64(sol) >= math.Pow(2, 31)-1 || float64(sol) <= -1*(math.Pow(2, 31)-1) {
		return 0
	}
	return sol
}
```

Solution in mind
================

-	Continually divide number by 10, keeping track of remainder.
-	Sum up these remainders by multiplying sum by 10 before adding each digit.
-	Finally check for overflow, if it occurred, return 0.
