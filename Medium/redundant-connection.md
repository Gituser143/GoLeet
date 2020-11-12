Code
====

```go
var seen map[int]bool
var graph map[int][]int

func dfs(source, target int) bool {
	if _, ok := seen[source]; !ok {
		seen[source] = true
		if source == target {
			return true
		}

		for _, neighbour := range graph[source] {
			if neighbour == target {
				return true
			}

			if dfs(neighbour, target) {
				return true
			}
		}
	}
	return false
}

func findRedundantConnection(edges [][]int) []int {
	graph = make(map[int][]int, 0)

	for _, nodes := range edges {
		u, v := nodes[0], nodes[1]
		seen = make(map[int]bool, 0)

		_, ok1 := graph[u]
		_, ok2 := graph[v]
		if ok1 && ok2 && dfs(u, v) {
			return nodes
		}

		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	return []int{}
}
```

Solution in mind
================

-	We create an empty seen set, and a graph. We begin by iterating through edges.

-	If the nodes of edge already exist in the graph and they can reach each other (check through dfs), that edge is duplicate.

-	Else add both nodes to the graph as vertices and neighbours.

-	The intuition is as we construct the tree (graph), if we encounter a cycle, the tree cannot be constructed, hence the edge we're trying to add is the duplicate edge.
