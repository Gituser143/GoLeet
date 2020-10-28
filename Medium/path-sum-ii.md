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

func isPathSum(root *TreeNode, sum int, nodes []int, solution [][]int) [][]int {
	if root == nil {
		return solution
	}

	nodes = append(nodes, root.Val)

	if sum-root.Val == 0 && root.Right == nil && root.Left == nil {
		path := make([]int, len(nodes))
		copy(path, nodes)
		solution = append(solution, path)
		return solution
	}

	solution = isPathSum(root.Left, sum-root.Val, nodes, solution)

	solution = isPathSum(root.Right, sum-root.Val, nodes, solution)

	return solution
}

func pathSum(root *TreeNode, sum int) [][]int {
	solution := isPathSum(root, sum, []int{}, [][]int{})
	return solution
}
```

Solution in mind
================

-	Approach similar to finding path, if a path is found, add path to solution and return solution.

-	To find path:

	-	If current node value - sum is equal to 0, AND the node is a leaf node, we have a valid path sum.

	-	If we ever encounter a nil node, we have reached the end of path and the path is not valid.

	-	We recursively call to see if either the right sub path or left sub path has a sum equal to sum - current node value.

Our recursive calls return solution slices (containing valid paths), this is returned at the end as solution.
