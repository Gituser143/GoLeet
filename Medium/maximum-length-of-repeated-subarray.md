Code
====

Naive Solution
--------------

```go
func findLength(nums1 []int, nums2 []int) int {
	maxLen := 0
	for i, num1 := range nums1 {
		if maxLen >= len(nums1)-i {
			return maxLen
		}
		for j, num2 := range nums2 {
			if num2 == num1 {
				count := 0
				for x := 0; i+x < len(nums1) && j+x < len(nums2); x++ {
					if nums1[i+x] == nums2[j+x] {
						count += 1
					} else {
						break
					}
				}

				if count > maxLen {
					maxLen = count
				}
			}
		}
	}

	return maxLen
}
```

With Dynamic Programming
------------------------

```go
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	maxLen := 0

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}
```

Solution in mind
================

Naive
-----

-	Here we simply iterate through possible combinations of `i` and `j` and we check if `nums1[i] == nums2[j]`.

-	If yes, we loop forward in both arrays till we get a non equal element.

-	A count (`count`) is maintained when looping forward through both arrays and is later updated as max count (`maxLen`).

Dynamic Programming
-------------------

-	We create an `m+1 * n+1` dp grid.

-	If ith element in nums1 is equal to jth element in nums2, we say `dp[i][j] = dp[i+1][j+1] + 1`.

-	Thus, if this is the first element being matched, it will be assigned as `0 + 1`. But if the element following it was also match, the match length increments.

-	We simultaneously keep track of max count seen and return this at the end.
