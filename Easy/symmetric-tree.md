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

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Right, node2.Left) && isMirror(node1.Left, node2.Right)
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}
```

Solution in mind
================

-	We recursively check if a tree is a mirror of itself.

-	We use an `isMirror(node1, node2)` function to check if the given two nodes are mirror copies.

-	for them to be a mirror copy:

	-	The roots should have the same value
	-	The right subtree of first node should be identical to left subtree of second.
	-	The left subtree of first node should be identical to right subtree of second.
