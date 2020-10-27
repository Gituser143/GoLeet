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

func postorder(root *TreeNode, traversal []int) []int {
	if root == nil {
		return traversal
	}

	traversal = postorder(root.Left, traversal)
	traversal = postorder(root.Right, traversal)
	traversal = append(traversal, root.Val)

	return traversal
}

func postorderTraversal(root *TreeNode) []int {
	traversal := postorder(root, []int{})
	return traversal
}
```
