Code
====

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

var parentNode *TreeNode

func recurseTree(node, p, q *TreeNode) bool {
	if node == nil {
		return false
	}

	l := recurseTree(node.Left, p, q)
	r := recurseTree(node.Right, p, q)

	mid := 0
	if node.Val == p.Val || node.Val == q.Val {
		mid = 1
	}

	left := 0
	if l {
		left = 1
	}

	right := 0
	if r {
		right = 1
	}

	if mid+right+left >= 2 {
		parentNode = node
	}

	return (mid + right + left) > 0
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentNode = nil
	recurseTree(root, p, q)

	return parentNode
}
```

Solution in mind
================

Solution, same as [lowest-common-ancestor-of-a-binary-tree](https://github.com/Gituser143/GoLeet/blob/main/Medium/lowest-common-ancestor-of-a-binary-tree.md)
