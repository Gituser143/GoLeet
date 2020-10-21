Code
====

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// Recursive function which traverses neigbhours of given node
func traverseNodes(node *Node, nodeMap map[int]*Node) map[int]*Node {
	if node != nil {
		// Iterate through neighbours
		for _, neighbour := range node.Neighbors {
			// if neighbour has not been visited, create clone and add to map
			// Then traverse the neighbour itself
			// Add cloned neighbour to list of neighbours to node
			if _, ok := nodeMap[neighbour.Val]; !ok {
				newNode := &Node{Val: neighbour.Val}
				nodeMap[neighbour.Val] = newNode
				nodeMap[node.Val].Neighbors = append(nodeMap[node.Val].Neighbors, newNode)
				nodeMap = traverseNodes(neighbour, nodeMap)
			} else {
				nodeMap[node.Val].Neighbors = append(nodeMap[node.Val].Neighbors, nodeMap[neighbour.Val])
			}
		}
	}
	return nodeMap
}

func cloneGraph(node *Node) *Node {
	nodeMap := make(map[int]*Node, 0)
	if node != nil {
		newnode := &Node{Val: node.Val}
		nodeMap[node.Val] = newnode
		nodeMap = traverseNodes(node, nodeMap)
	} else {
		return nil
	}
	return nodeMap[1]
}
```

Solution in mind
================

-	Create a map to store `Node value: Pointer to node`. Add initial node to map.
-	Traverse neighbours recursively. If new neighbours are found create clone of neighbour and add to list of cloned node's neighbours.
-	Traverse the newly found neighbour.
