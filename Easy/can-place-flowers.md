Code
====

```go
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	count := 0
	l := len(flowerbed)

	if l == 1 {
		if flowerbed[0] == 0 && n == 1 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < l; i++ {
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count += 1
			}
		}

		if i == l-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				count += 1
			}
		}

		if i > 0 && i < l && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count += 1
		}

		if count >= n {
			return true
		}

	}

	return false
}
```

Solution in mind
================

-	Iterate through flowerbed, if a position is empty and adjacent positions are also empty, fill the position and increment `count`.

-	If `count` equals or is greater than `n`, return true.

-	Return false if `count` not equal to `n` outside the loop (`n` plants cannot be plotted).
