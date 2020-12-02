Code
====

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	nums   []int
	length int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	l := 0
	nums := []int{}
	temp := head
	for temp != nil {
		nums = append(nums, temp.Val)
		temp = temp.Next
		l += 1
	}

	sol := Solution{nums, l}

	return sol
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	num := r.Intn(this.length)

	return this.nums[num]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
```

Solution in mind
================

-	Initialise a solution, with a slice of numbers in linked list and the length of the slice.

-	To generate a random node's value, generate an index and retrun the number in the solution slice at said index
