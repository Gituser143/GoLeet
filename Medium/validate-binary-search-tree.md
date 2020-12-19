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

func getInorder(root *TreeNode, order []int) []int {
	if root == nil {
		return order
	}

	order = getInorder(root.Left, order)

	order = append(order, root.Val)

	order = getInorder(root.Right, order)

	return order
}

func isOrdered(arr []int) bool {
	n := len(arr)

	if n < 2 {
		return true
	}

	for i := 1; i < n; i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	arr := getInorder(root, []int{})

	return isOrdered(arr)
}
```

Solution in mind
================

-	Perform an inorder traversal of BST, keeping track of values in an array. An inorder traversal of a valid BST should give a sorted (strictly increasing) list of numbers.

-	Iterate through the array, if an element is lesser than or equal to it's previous, the array is not strictly increasing and thus the tree is not a BST.
