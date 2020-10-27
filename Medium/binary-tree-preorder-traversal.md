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

func preorder(root *TreeNode, traversal []int) []int {
	if root == nil {
		return traversal
	}

	traversal = append(traversal, root.Val)
	traversal = preorder(root.Left, traversal)
	traversal = preorder(root.Right, traversal)

	return traversal
}

func preorderTraversal(root *TreeNode) []int {
	traversal := preorder(root, []int{})
	return traversal
}
```
