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

func traverse(root *TreeNode, low int, high int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= low && root.Val <= high {
		*sum += root.Val
	}

	traverse(root.Right, low, high, sum)
	traverse(root.Left, low, high, sum)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	traverse(root, low, high, &sum)
	return sum
}
```

Solution in mind
================

-	Perform a traversal on the tree (postorder, preorder or inorder).

-	On encountering a node with a value within the range, increment the value of the sum pointer.

-	Return value of sum pointer after traversal.
