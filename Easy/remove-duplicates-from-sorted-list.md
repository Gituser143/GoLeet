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
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	temp := curr

	for temp != nil {
		for temp != nil && temp.Val == curr.Val {
			temp = temp.Next
		}
		curr.Next = temp
		curr = temp
	}

	return head
}
```

Solution in mind
================

-	Make use of a temp pointer to skip duplicate nodes.

-	After skipping of nodes, link the previous pointer to it's new node.
