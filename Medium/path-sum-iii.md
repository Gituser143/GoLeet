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

func getAllPaths(root *TreeNode, sum int, nodes []int, solution int) int {
	if root == nil {
		return solution
	}

	nodes = append(nodes, root.Val)

	pathSum := 0
	for i := len(nodes) - 1; i >= 0; i-- {
		pathSum += nodes[i]
		if pathSum == sum {
			solution++
		}
	}

	solution = getAllPaths(root.Left, sum, nodes, solution)
	solution = getAllPaths(root.Right, sum, nodes, solution)

	return solution
}

func pathSum(root *TreeNode, sum int) int {
	allPaths := getAllPaths(root, sum, []int{}, 0)
	return allPaths
}
```

Solution in mind
================

-	For every encountered node, check if sum of node and parents upto root equals the target sum.

-	If yes, increment a count variable to keep track of valid paths.

-	Recursively call the function on children and return solution.
