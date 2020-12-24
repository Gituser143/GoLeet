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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	next := head.Next

	head = next
	prev := head

	curr.Next = next.Next
	next.Next = curr
	prev = curr
	curr = curr.Next
	if curr == nil {
		return head
	}

	next = curr.Next

	for next != nil {
		prev.Next = next
		curr.Next = next.Next
		next.Next = curr

		prev = curr
		curr = curr.Next
		if curr == nil {
			break
		}
		next = curr.Next
	}

	return head
}
```

Solution in mind
================

-	To swap nodes in pairs, we make use of 3 pointers: `prev`, `curr`, `next`, which are pointers to the previous node, current node and next node.

-	In general, two nodes are swapped when the current node `curr`, and the node after it `next` modify their `Next` pointers. The current node `curr` points to the node after next, `next.Next` and `next` points to `curr`. This swaps their positions.

-	But a key point to note here is that the node before `curr` (Before swapping), must now point to the new node (Swapped `next`). Thus after swapping, we make `prev` point to `next`.

-	After the links have been swapped, we move each pointer to their next places. These are: `prev` to `curr` (to the new previous node), `curr` to `curr.Next` (to the next node to be swapped), `next` to `curr.Next` (node after `curr`).

-	We repeat the swapping process in a loop till either `next` or `curr` turns into `nil`. Although, this loop is only for after swapping the first two nodes. This is because there exists no `prev` for the first two nodes (Swapped explicitly).
