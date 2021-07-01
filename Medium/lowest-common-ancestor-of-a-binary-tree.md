Code
====

```go
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

-	Recursively search for nodes

-	If the nodes are in the left subtree and the right subtree of the current node, the current node is the solution.

-	Additionally, if a node is in either one of the subtrees and the other node is the current node itself, the current node is the solution.

Naive Code
==========

```go
func getParents(root, node *TreeNode, parents []*TreeNode) []*TreeNode {
	if root == nil {
		return parents
	}

	if root.Val == node.Val {
		parents = append([]*TreeNode{root}, parents...)
		return parents
	}

	parents = getParents(root.Left, node, parents)
	if len(parents) != 0 {
		parents = append([]*TreeNode{root}, parents...)
		return parents
	}

	parents = getParents(root.Right, node, parents)
	if len(parents) != 0 {
		parents = append([]*TreeNode{root}, parents...)
		return parents
	}

	return parents
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pParents := getParents(root, p, []*TreeNode{})
	qParents := getParents(root, q, []*TreeNode{})

	i := 0

	for pParents[i].Val == qParents[i].Val {
		i += 1
		if i >= len(pParents) || i >= len(qParents) {
			break
		}
	}

	return pParents[i-1]
}
```

Solution in mind
================

-	Get the parents for both nodes.

-	Iterate through parents and the last matching parent is the solution.
