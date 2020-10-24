Code
====

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var prevElement *TreeNode
var firstElement *TreeNode
var secondElement *TreeNode

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	traverse(root.Left)

	if firstElement == nil && (prevElement == nil || prevElement.Val > root.Val) {
		firstElement = prevElement
	}

	if firstElement != nil && prevElement.Val > root.Val {
		secondElement = root
	}

	prevElement = root

	traverse(root.Right)
}

func recoverTree(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	prevElement = nil
	firstElement = nil
	secondElement = nil

	traverse(root)

	firstElement.Val, secondElement.Val = secondElement.Val, firstElement.Val
}
```

Solution in mind
================

-	When traversing the tree in inorder, nodes are in ascending order.

-	We have to find the two nodes that are out of place.

-	The first is greater than the next element while the second is less than the previous.

-	We make use of a recursive inorder function to find these elements. We keep global pointers to keep track of previous element, first anomaly and second anomaly.

-	After the traversal, we swap the values of the 2 anomalies.
