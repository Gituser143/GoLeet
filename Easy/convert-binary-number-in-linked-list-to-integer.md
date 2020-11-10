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

var total float64
var i float64

func traverse(head *ListNode) {
	if head == nil {
		return
	}

	traverse(head.Next)
	if head.Val != 0 {
		total += math.Pow(2, i)
	}

	i++

}

func getDecimalValue(head *ListNode) int {
	total = 0
	i = 0

	traverse(head)

	return int(total)
}
```

Solution in mind
================

-	Iterate through the linked list in reverse and summate non zero indices as growing powers of 2.
