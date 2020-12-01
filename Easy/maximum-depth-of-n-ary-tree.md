Code
====

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	depth := 0
	if root == nil {
		return depth
	}

	level := []*Node{root}

	for len(level) != 0 {
		newLevel := []*Node{}
		depth += 1
		for _, node := range level {
			for _, child := range node.Children {
				if child != nil {
					newLevel = append(newLevel, child)
				}
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
