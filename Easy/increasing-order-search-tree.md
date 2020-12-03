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

func inOrder(root *TreeNode, vals []int) []int {
	if root == nil {
		return vals
	}
	vals = inOrder(root.Left, vals)
	vals = append(vals, root.Val)
	vals = inOrder(root.Right, vals)

	return vals
}

func increasingBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	vals := inOrder(root, []int{})

	solNode := &TreeNode{Val: vals[0]}
	ptrNode := solNode
	for _, val := range vals[1:] {
		newNode := &TreeNode{Val: val}
		ptrNode.Right = newNode
		ptrNode = ptrNode.Right
	}

	return solNode
}
```

Solution in mind
================

-	Perform an inorder traversal of BST to get all values in a slice. Inorder traversal of BST gives sorted vals.

-	Construct new BST from values and return root node.
