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

func reverse(start, end *ListNode) (*ListNode, *ListNode) {
	curr := start
	var prev *ListNode = nil
	var next *ListNode = nil
	for curr != end {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev

	return end, start
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	start := head
	var prev *ListNode = nil
	for i := 1; i < left; i++ {
		prev = start
		start = start.Next
	}

	end := head
	for i := 1; i < right; i++ {
		end = end.Next
	}
	next := end.Next

	reverse(start, end)

	if prev != nil {
		prev.Next = end
	}
	start.Next = next

	if left == 1 {
		head = end
	}

	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	n := 0
	temp := head

	for temp != nil {
		temp = temp.Next
		n += 1
	}

	for i := 1; i < n; i += k {
		j := i + k - 1
		if j > n {
			break
		}
		head = reverseBetween(head, i, j)
	}

	return head
}
```

Solution in mind
================

-	The solution builds upon the problem [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/). The solution for the same can be found [here](https://github.com/Gituser143/GoLeet/blob/main/Medium/reverse-linked-list-ii.md)

-	We first get a count of existing nodes as `n`.

-	We then iterate from `i = 0 to n`, in intervals of `k`, while performing `reverseBetween()` on `i, i+k-1`.
