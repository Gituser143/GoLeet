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
		mergedSol := &ListNode{}
		temp := mergedSol
		for lists[i] != nil && solutionList != nil {
			if lists[i].Val > solutionList.Val {
				mergedSol.Next = solutionList
				solutionList = solutionList.Next
				mergedSol = mergedSol.Next
			} else {
				mergedSol.Next = lists[i]
				lists[i] = lists[i].Next
				mergedSol = mergedSol.Next
			}
		}

		if lists[i] == nil {
			mergedSol.Next = solutionList
		} else if solutionList == nil {
			mergedSol.Next = lists[i]
		}

		solutionList = temp.Next
	}

	return solutionList
}
```

Solution in mind
================

-	Set the first list as the solution list.

-	Go through remaining lists, merging each with the solution list.

-	Merge is done as:

	-	While lists are non empty.
	-	Compare heads of both lists.
	-	Make merged list point to lower.
	-	Increment node pointer of lower.
	-	Move merge pointer to new node.
