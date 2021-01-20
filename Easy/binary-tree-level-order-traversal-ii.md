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
func levelOrderBottom(root *TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}

	levelNodes := []*TreeNode{root}

	for len(levelNodes) != 0 {
		level := []int{}
		nextLevel := []*TreeNode{}
		for _, node := range levelNodes {
			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		levelNodes = nextLevel
		levels = append(levels, []int{})
		copy(levels[1:], levels[:len(levels)-1])
		levels[0] = level
	}

	return levels
}
```
