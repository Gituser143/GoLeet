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
func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	level := []*TreeNode{root}

	for len(level) != 0 {
		newLevel := []*TreeNode{}
		depth += 1
		for _, node := range level {
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}

			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
		}

		level = newLevel
	}

	return depth
}
```

Solution in mind
================

-	Perform a level order traversal using a queue.

-	Increment depth on reaching new level. Compute next levels as children of current level.

-	On encountering a leaf node, return depth as first leaf node encountered in level order will be of minimum depth.
