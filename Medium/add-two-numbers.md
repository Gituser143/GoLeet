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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	current := &head
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				val := current.Val + l2.Val
				current.Val = val % 10
				carry = val / 10

				if l2.Next != nil || carry != 0 {
					nextNode := ListNode{carry, nil}
					current.Next = &nextNode
					current = current.Next
				}

				l2 = l2.Next
			}
		} else if l2 == nil {
			for l1 != nil {
				val := current.Val + l1.Val
				current.Val = val % 10
				carry = val / 10

				if l1.Next != nil || carry != 0 {
					nextNode := ListNode{carry, nil}
					current.Next = &nextNode
					current = current.Next
				}

				l1 = l1.Next
			}
		} else {
			val := current.Val + l1.Val + l2.Val
			current.Val = (val % 10)
			carry = val / 10

			if l1.Next != nil || l2.Next != nil || carry != 0 {
				nextNode := ListNode{carry, nil}
				current.Next = &nextNode
				current = current.Next
			}

			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return &head
}
```

Solution in mind
================

-	Create new linked list to store solution.
-	Iterate through both linked lists.
-	Sum numbers of linked list and store units place in new node.
-	Store carry in next node.
-	Repeat till either linked list is nil. Then add last carry + numbers of remaining linked list.
