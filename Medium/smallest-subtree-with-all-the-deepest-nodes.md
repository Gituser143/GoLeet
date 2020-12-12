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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	parents := make(map[*TreeNode]*TreeNode, 0)
	level := []*TreeNode{root}

	for len(level) != 0 {
		nextLevel := []*TreeNode{}
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
				parents[node.Left] = node
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
				parents[node.Right] = node
			}
		}

		if len(nextLevel) == 0 {
			break
		}

		level = nextLevel
	}

	if len(level) == 1 {
		return level[0]
	}

	for len(level) > 0 {
		nextParents := make(map[*TreeNode]bool, 0)
		nextLevel := []*TreeNode{}
		for _, node := range level {
			if _, ok := nextParents[parents[node]]; !ok {
				nextParents[parents[node]] = true
				nextLevel = append(nextLevel, parents[node])
			}
		}

		if len(nextLevel) == 1 {
			return nextLevel[0]
		}

		level = nextLevel
	}

	return root
}
```

Solution in mind
================

-	Traverse the tree in level order and maintain 2 data structures `level` (list of nodes in current level) and `parents` (hashMap of node: parent).

-	After traversal, `level` will have the nodes in the deepest level.

-	If only one node exists in this list, return the node as it will be the deepest subtree.

-	If more than one exist, iterate through each node in the `level` and maintain a set `nextParents` (keeps track of parents) and `nextLevel` (keeps track of parents of deepest nodes). We move upwards till we find a parent that has all deepest nodes as children.

-	Update `level` as `nextLevel` once parsed fully. When `level` has only 1 node left in it, that node is root of the deepest subtree. We then return that node.
