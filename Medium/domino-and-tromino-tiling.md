Code
====

```go
func numTilingRec(n int, dp *[]int) int {
	if n <= 3 {
		return (*dp)[n]
	}

	res := 0
	if (*dp)[n-1] != -1 {
		res += 2 * (*dp)[n-1]
	} else {
		res += (2 * numTilingRec(n-1, dp)) % 1000000007
	}

	if (*dp)[n-3] != -1 {
		res += (*dp)[n-3]
	} else {
		res += (numTilingRec(n-3, dp)) % 1000000007
	}

	(*dp)[n] = res % 1000000007
	return (*dp)[n]
}

func numTilings(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	return numTilingRec(n, &dp)
}
```

Solution in mind
================

We can derive the possible ways to place tiles at a given width `n` by the below formula:

```
dp[n] = dp[n-1] + dp[n-2] + 2*(dp[n-3] + ... + dp[0])
```

The above formula is derived as, for an `n` length tile, we can get the combinations as:

-	`n-1` combinations + 1 vertical tile
-	`n-2` combinations + 2 horizontal tiles
-	For `n-3` upto `n-n`, we have 2 times the respective combinations. This is achieved by placing 2 trominos at the ends and filling the middle with horizontal dominoes. The 2 trominos can be rotated, thus forming 2 sets of combinations.

We can further simply the above formula as follows:

```
dp[n] = dp[n-1] + dp[n-2] + 2*(dp[n-3] + ... + dp[0])      =>   (i)
dp[n-1] = dp[n-2] + dp[n-3] + 2*(dp[n-4] + ... + dp[0])    =>   (ii)

```

Subtracting `ii` from `i`, we get:

```
dp[n] - dp[n-1] = dp[n-1] + dp[n-3]
```

Which is simplified as:

```
dp[n] = 2*dp[n-1] + dp[n-3]
```

Using the above equation, we can perform recursive calls to calculate the possible combinations for width `n`.
