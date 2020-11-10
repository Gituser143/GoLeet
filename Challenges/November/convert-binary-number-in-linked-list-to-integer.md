Statement
=========

<p>Given <code>head</code> which is a reference node to&nbsp;a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.</p>

<p>Return the <em>decimal value</em> of the number in the linked list.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2019/12/05/graph-1.png" style="width: 426px; height: 108px;">

```
Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10
```
<p><strong>Example 2:</strong></p>

```
Input: head = [0]
Output: 0
```

<p><strong>Example 3:</strong></p>

```
Input: head = [1]
Output: 1
```

<p><strong>Example 4:</strong></p>

```
Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
Output: 18880
```

<p><strong>Example 5:</strong></p>

```
Input: head = [0,0]
Output: 0
```

<p><strong>Constraints:</strong></p>

<ul>
    <li>The Linked List is not empty.</li>
    <li>Number of nodes&nbsp;will not exceed <code>30</code>.</li>
    <li>Each node's value is either&nbsp;<code>0</code> or <code>1</code>.</li>
</ul>

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
