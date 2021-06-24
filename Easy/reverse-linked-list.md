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

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	var next *ListNode = nil

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head = prev

	return head
}
```

Solution in mind
================

-	Keep track of 3 pointer, `prev`, `curr` and `next`. With `curr` being initialised to `head` and the others being `nil`.

-	Move through `curr` till we reach the end of linked list.

-	Set `next` as `curr.Next`. Now the current node's pointer can be set to the previous node.

-	After updating the current node's pointer, move `curr` to `next` and `prev` to `curr`.

-	Finally, set the new `head` as `prev`, this will update the head to the new beginning of the list.
