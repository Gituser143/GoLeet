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
func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	level := []*TreeNode{root}

	for len(level) != 0 {
		newLevel := []*TreeNode{}
		depth += 1
		for _, node := range level {
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

-	Perform a simple level order traversal using a queue.

-	Iterate through each level, populating next level using non nil children of current level.

-	For each level traversed, increment depth.
