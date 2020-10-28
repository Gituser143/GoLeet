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

func getAllPaths(root *TreeNode, nodes string, solution []string) []string {
	if root == nil {
		return solution
	}

	nodes = nodes + fmt.Sprint(root.Val) + "->"

	if root.Right == nil && root.Left == nil {
		solution = append(solution, nodes[:len(nodes)-2])
		return solution
	}

	solution = getAllPaths(root.Left, nodes, solution)
	solution = getAllPaths(root.Right, nodes, solution)

	return solution
}

func binaryTreePaths(root *TreeNode) []string {
	solution := getAllPaths(root, "", []string{})
	return solution
}
```

Solution in mind
================

-	Recursively call on left sub tree and right sub tree

-	Keep track of traversed path and append node values to path.

-	If leaf node, append path to solution list.

-	Return solution list
