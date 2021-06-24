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
```

Solution in mind
================

-	Given a `left` and `right` index, we can get the `start` and `end` nodes. All nodes in between these (inclusive of `start` and `end`) must be reversed. So we pass them as arguments to `reverse()`.

-	The function `reverse()` will reverse a list given the starting and end node.

-	Post reversal, if `start` is not the first node, the node preceding it (before reversal) must be linked to `end`. Similarly, the node after end (before reversal) must be linked to `start`.

-	If `start` is the first node (`left == 1`), then `head` must be updated to point to the new start of the linked list, i.e, `end`.
