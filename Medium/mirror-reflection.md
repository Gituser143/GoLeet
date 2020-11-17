Code
====

```go
func gcd(a, b int) int {
	for a > 0 && b > 0 {
		if a == b {
			return a
		} else if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a == 0 {
		return b
	} else {
		return a
	}

	return 0
}

func mirrorReflection(p int, q int) int {
	g := gcd(p, q)

	p /= g
	p %= 2

	q /= g
	q %= 2

	if p == 1 && q == 1 {
		return 1
	}

	if p != 1 {
		return 2
	}

	return 0

}
```

Solution in mind
================

-	Implemented Mathematical solution as given [here](https://leetcode.com/problems/mirror-reflection/solution/)

-	Instead of modelling the ray as a bouncing line, model it as a straight line through reflections of the room.

-	For example, if `p = 2, q = 1`, then we can reflect the room horizontally, and draw a straight line from `(0, 0)` to `(4, 2)`. The ray meets the receptor `2`, which was reflected from `(0, 2)` to `(4, 2)`.

-	In general, the ray goes to the first integer point `(kp, kq)` where `k` is an integer, and `kp` and `kq` are multiples of `p`. Thus, the goal is just to find the smallest `k` for which `kq` is a multiple of `p`.

-	The mathematical answer is `k = p / gcd(p, q)`.
