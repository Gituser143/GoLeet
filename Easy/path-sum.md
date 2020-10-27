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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val-sum == 0 {
		if root.Right == nil && root.Left == nil {
			return true
		}
	}

	return hasPathSum(root.Right, sum-root.Val) || hasPathSum(root.Left, sum-root.Val)
}
```

Solution in mind
================

-	If current node value - sum is equal to 0, AND the node is a leaf node, we have a valid path sum.

-	If we ever encounter a nil node, we have reached the end of path and the path is not valid.

-	We recursively call to see if either the right sub path or left sub path has a sum equal to sum - current node value.
