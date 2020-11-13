Code
====

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	level := []*Node{root}

	for len(level) > 0 {
		nextLevel := []*Node{}
		i := 0
		for i = 0; i < len(level)-1; i++ {
			level[i].Next = level[i+1]
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}

			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}

		level[i].Next = nil

		if level[i].Left != nil {
			nextLevel = append(nextLevel, level[i].Left)
		}

		if level[i].Right != nil {
			nextLevel = append(nextLevel, level[i].Right)
		}

		level = nextLevel
	}

	return root
}
```

Solution in mind
================

### Note: Solution is almost ditto to populating-next-right-pointers-in-each-node

-	To solve this, iterate over nodes level by level.

-	We store all nodes in a level in a queue. Starting with just the root node.

-	We iterate through this queue and set the pointer of each node to the next node in queue.

-	While doing this, we add children of nodes we've parsed into a temporary queue.

-	Once the original queue has been iterated through, we replace it with the temp queue, thus replacing one level by the level below it.
