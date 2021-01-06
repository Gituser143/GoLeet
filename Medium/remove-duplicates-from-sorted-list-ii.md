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

	if head == nil || head.Next == nil {
		return head
	}

	sentinal := &ListNode{-1, head}

	prev := sentinal
	curr := head

	for curr != nil {

		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
		} else {
			prev = prev.Next
		}

		curr = curr.Next
	}

	return sentinal.Next
}
```

Solution in mind
================

-	Solution as described [here](https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/solution/)
