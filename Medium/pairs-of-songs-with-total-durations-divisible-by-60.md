Code
====

```go
func numPairsDivisibleBy60(time []int) int {
	pairs := 0

	remainder := make(map[int]int, 0)

	for _, song := range time {
		rem := song % 60
		if _, ok := remainder[rem]; ok {
			remainder[rem] += 1
		} else {
			remainder[rem] = 1
		}
	}

	for rem1, s1 := range remainder {
		if rem1 == 0 || rem1 == 30 {
			n := remainder[rem1]
			pairs += n * (n - 1) / 2
			delete(remainder, rem1)
			continue
		}

		if _, ok := remainder[60-rem1]; ok {
			n := s1
			m := remainder[60-rem1]

			pairs += n * m

			delete(remainder, rem1)
		}

	}

	return pairs
}
```

Solution in mind
================

-	To get pairs of songs which duration as multiple of 60, we maintain a hashMap (`remainder`), indexed by the value `duration % 60`, holding count of all songs of said remainder value.

-	We construct the map by iterating through the array `time` and incrementing/adding values into `remainder`.

-	Once the `remainder` map is populated, we iterate through remainder values (`rem1`) and check if its compliment (`60-remainder`) exists in the map. If it does, we can form `n * m` pairs, `n` being the number of elements with remainder `rem1` and `m` being number of elements with remainder `60-rem1`.

-	In special cases of the remainder being either 30 (60/2) or 0, we can form `n * (n - 1) / 2` pairs, as these numbers pair only amongst themselves.
