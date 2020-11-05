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

var nodeIndex int
var currentIndex int

func traverseNode(node *ListNode) {
	if node == nil {
		currentIndex++
		return
	}

	traverseNode(node.Next)

	if currentIndex > nodeIndex {
		node.Next = node.Next.Next
		currentIndex = -1
		return
	}

	if currentIndex == -1 {
		return
	}

	currentIndex++
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	currentIndex = 0
	nodeIndex = n

	traverseNode(head)

	if currentIndex > n {
		for currentIndex > n && head != nil {
			head = head.Next
			currentIndex--
		}
	}

	return head
}
```

Solution in mind
================

-	Recursively traverse the linked list.

-	On encountering `nil`, increment a counter and return.

-	After recursive call to traverse next node, check if the counter is greater than n. If yes, the next node is to be removed. Set the counter to -1 and return.

-	We check if the counter is negative to avoid checking as we would have found an element.

-	If counter is lesser, increment it and return.

-	To handle cases like removal of first element or only element in list, we check if counter is still greater than n after traversal. If it is, we move head to next till counter is not greater than n or head is nil.

-	The required solution is now in head.
