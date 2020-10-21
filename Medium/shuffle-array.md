Code
====

```go
import "math/rand"

type Solution struct {
	nums     []int
	original []int
}

// Initialise the Solution struct
func Constructor(nums []int) Solution {
	original := []int{}
	for i, _ := range nums {
		original = append(original, nums[i])
	}
	numbers := make([]int, len(original))
	sol := Solution{nums: numbers, original: original}
	return sol
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	// Seed the random function
	rand.Seed(time.Now().UnixNano())

	// Copy original into nums
	// Always shuffle the original array
	copy(this.nums, this.original)

	// Shuffle by swapping random elements
	rand.Shuffle(len(this.original), func(i, j int) {
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	})
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
```
