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

func pruneRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return root.Val
	}

	lSum := pruneRecursive(root.Left)
	if lSum == 0 {
		root.Left = nil
	}

	rSum := pruneRecursive(root.Right)
	if rSum == 0 {
		root.Right = nil
	}

	return rSum + lSum + root.Val
}

func pruneTree(root *TreeNode) *TreeNode {
	sum := pruneRecursive(root)
	if sum == 0 {
		return nil
	}

	return root
}
```

Solution in mind
================

-	We calculate sums of left and right subtree. If either are 0, then there were no 1's in that subtree and it can be pruned (set to `nil`).

-	Sum is calculated recursively by using `node.Val` + `lSum` + `rSum`.
