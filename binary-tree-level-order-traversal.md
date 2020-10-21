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

func levelOrder(root *TreeNode) [][]int {
	// Init empty levels
	levels := [][]int{}

	// Return for empty tree
	if root == nil {
		return levels
	}

	// Initialise slice to hold nodes at a level
	// Init to root
	tempLevels := [][]*TreeNode{[]*TreeNode{root}}

	// While there are nodes in the level
	for len(tempLevels) != 0 {
		// nextTempLevel holds nodes in children of tempLevels[0]
		// nextLevel holds values of nodes in current level
		nextTempLevel := []*TreeNode{}
		nextLevel := []int{}

		// Iterate over nodes in current level
		for _, node := range tempLevels[0] {
			// Append children to next level
			if node != nil && node.Left != nil {
				nextTempLevel = append(nextTempLevel, node.Left)
			}

			if node != nil && node.Right != nil {
				nextTempLevel = append(nextTempLevel, node.Right)
			}

			// Append current node values to nextLevel
			if node != nil {
				nextLevel = append(nextLevel, node.Val)
			}
		}

		// Append current node values to solution slice
		levels = append(levels, nextLevel)

		// Append next level to parse if non empty
		if len(nextTempLevel) > 0 {
			tempLevels = append(tempLevels, nextTempLevel)
		}

		// Drop currently parsed level
		tempLevels = tempLevels[1:]
	}

	return levels
}
```

Solution in mind
================

-	Start at root, create slice of children.
-	Iterate through children slice and add children of each iterated node to a new slice.
-	While adding children of node to new slice, store parent node values into solution.
-	Replace current slice with children slice and repeat.
