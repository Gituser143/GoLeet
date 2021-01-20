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
func averageOfLevels(root *TreeNode) []float64 {
	levels := []float64{}
	if root == nil {
		return levels
	}

	levelNodes := []*TreeNode{root}

	for len(levelNodes) != 0 {
		sum := 0
		count := 0
		nextLevel := []*TreeNode{}
		for _, node := range levelNodes {
			sum += node.Val
			count += 1
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		levelNodes = nextLevel
		levels = append(levels, float64(sum)/float64(count))
	}

	return levels
}
```

Solution in mind
================

-	Perform a level order traversal.

-	Keep track of sum and number of nodes at every level.

-	Store sum/count for every level
