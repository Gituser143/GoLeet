Code
====

```go
func trap(height []int) int {
	left := 0
	right := len(height) - 1

	ans := 0

	leftMax := 0
	rightMax := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}

	return ans
}
```

Solution in mind
================

-	We use two pointers, left and right. We keep track of left highest value and right highest.

-	We find the lower of left and right heights, if they are lower than their respective max values, we add max - value to the answer.

-	If they are greater, we update the max values.

-	We then move the lower pointer by one position towards the other pointer.

-	We repeat the above steps till the left and right pointers meet.

![Animation](https://leetcode.com/problems/trapping-rain-water/solution/)
