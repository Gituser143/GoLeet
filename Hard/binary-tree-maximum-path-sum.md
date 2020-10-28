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

func getMaxPath(root *TreeNode, maxPathVal *int) int {
	if root == nil {
		return 0
	}

	// Get max path of left sub tree
	l := getMaxPath(root.Left, maxPathVal)
	if l < 0 {
		l = 0
	}

	// Get max path of right sub tree
	r := getMaxPath(root.Right, maxPathVal)
	if r < 0 {
		r = 0
	}

	// Update the global maxPathVal if left path to right path is greater
	if root.Val+l+r > *maxPathVal {
		*maxPathVal = root.Val + l + r
	}

	// Find greater of two paths
	maxPath := 0
	if l > r {
		maxPath = l
	} else {
		maxPath = r
	}

	// Return highest value path upto root
	return root.Val + maxPath
}

func maxPathSum(root *TreeNode) int {
	maxPathVal := -1000
	getMaxPath(root, &maxPathVal)
	return maxPathVal
}
```

Solution in mind
================

-	Keep track of max path value in global variable (Init to min node value).

-	Recursively find max path of left sub tree and right sub tree.

-	If right path + left path + root > max path, update max path.

-	Return max(left path, right path) + root value as max path of sub tree.

-	After traversal, global variable will have max path value.
