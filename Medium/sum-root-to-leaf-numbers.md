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

func getAllNumbers(root *TreeNode, path int, solution []int) []int {
	if root == nil {
		return solution
	}

	path = (10 * path) + root.Val

	if root.Right == nil && root.Left == nil {
		solution = append(solution, path)
		return solution
	}

	solution = getAllNumbers(root.Left, path, solution)
	solution = getAllNumbers(root.Right, path, solution)

	return solution
}

func sumNumbers(root *TreeNode) int {
	solution := getAllNumbers(root, 0, []int{})
	sum := 0
	for _, val := range solution {
		sum += val
	}
	return sum
}
```

Solution in mind
================

-	Recursively call a function on left sub and right subtree to generate a path.

-	If node is not nil, add value of node to path.

-	If a leaf node is encountered, add path to solution list.

-	Return solution list
