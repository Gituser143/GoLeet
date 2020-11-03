Code
====

```go
func findMinHeightTrees(n int, edges [][]int) []int {
	if n <= 2 {
		solution := []int{}
		for i := 0; i < n; i++ {
			solution = append(solution, i)
		}
		return solution
	}

	graph := make(map[int][]int, 0)
	orders := make(map[int]int, 0)

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]

		if _, ok := graph[node1]; ok {
			graph[node1] = append(graph[node1], node2)
			orders[node1] += 1
		} else {
			graph[node1] = []int{node2}
			orders[node1] = 1
		}

		if _, ok := graph[node2]; ok {
			graph[node2] = append(graph[node2], node1)
			orders[node2] += 1
		} else {
			graph[node2] = []int{node1}
			orders[node2] = 1
		}
	}

	remainingNodes := n
	leaves := []int{}

	for node, order := range orders {
		if order == 1 {
			leaves = append(leaves, node)
		}
	}

	for remainingNodes > 2 {
		remainingNodes -= len(leaves)
		newLeaves := []int{}

		for len(leaves) > 0 {
			leafNode := leaves[0]
			leaves = leaves[1:]
			for _, neighbour := range graph[leafNode] {
				orders[neighbour] -= 1
				if orders[neighbour] == 1 {
					newLeaves = append(newLeaves, neighbour)
				}
			}
		}

		leaves = append(leaves, newLeaves...)
	}

	return leaves
}
```

Solution in mind
================

-	The required minimum height trees will have roots which are the centroids of the graph.

-	These centroids are either 1 or 2 in number.

-	We can find these centroids by running a variation of topological sort:

	-	We maintain a queue of leaf nodes.
	-	We iteratively remove these leaf nodes from the queue.
	-	As we remove leaf nodes, more leaf nodes are created, we add these to the queue.
	-	We repeat the leaf removal process till the nodes remaining are greater than 2. (Max 2 centroids in a graph)

-	After the repeated leaf removal process, layers of the graph are removed through each iteration till only the centroids remain. We then return the list of centroids as our solution.
