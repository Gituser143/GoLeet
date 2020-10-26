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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
			head = head.Next
		} else {
			head.Next = l2
			l2 = l2.Next
			head = head.Next
		}
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}

	return temp.Next
}
```

Solution in mind
================

-	While lists are non empty, compare heads of both lists.
-	Make merged list point to lower.
-	Increment node pointer of lower.
-	Move merge pointer to new node.
-	If either list is empty and other is not, make merged pointer point to beginning of list.
