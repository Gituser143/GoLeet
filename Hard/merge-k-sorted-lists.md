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
func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	solutionList := lists[0]

	for i := 1; i < len(lists); i++ {

		for lists[i] != nil {
			listNode := lists[i]
			var prev *ListNode = nil
			solutionNode := solutionList

			for solutionNode != nil && solutionNode.Val < listNode.Val {
				prev = solutionNode
				solutionNode = solutionNode.Next
			}

			lists[i] = lists[i].Next

			if prev == nil {
				listNode.Next = solutionNode
				solutionList = listNode
			} else {
				prev.Next = listNode
				listNode.Next = solutionNode
			}
		}
	}

	return solutionList
}
```

Solution in mind
================

-	Set the first list as the solution list.

-	Go through remaining lists, merging each with the solution list.

-	Merge is done as:

	-	Iterate through each node in list
	-	Find appropriate position of node in solution list
	-	Move node to solution list.
