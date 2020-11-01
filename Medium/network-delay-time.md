Code
====

```go
func dfs(node int, elapsed int, graph map[int][][]int, nodeDistances map[int]int) map[int]int {
	if elapsed >= nodeDistances[node] {
		return nodeDistances
	}

	nodeDistances[node] = elapsed

	for _, val := range graph[node] {
		neighbour, time := val[0], val[1]
		nodeDistances = dfs(neighbour, elapsed+time, graph, nodeDistances)
	}

	return nodeDistances
}

func networkDelayTime(times [][]int, N int, K int) int {
	// node: [[neighbour, time], [nei, time]]
	graph := make(map[int][][]int, 0)

	for _, time := range times {
		if _, ok := graph[time[0]]; ok {
			graph[time[0]] = append(graph[time[0]], []int{time[1], time[2]})
		} else {
			graph[time[0]] = [][]int{[]int{time[1], time[2]}}
		}
	}

	// node: distance to node from src (K)
	nodeDistances := make(map[int]int, 0)
	for i := 1; i < N+1; i++ {
		nodeDistances[i] = 100 * 101
	}

	minPath := 0

	nodeDistances = dfs(K, 0, graph, nodeDistances)

	for _, val := range nodeDistances {
		if val > minPath {
			minPath = val
		}
	}

	if minPath == 100*101 {
		return -1
	} else {
		return minPath
	}

	return 0
}
```

Solution in mind
================

-	We first create a graph of format node: `[[neighbour, time], [nei, time]]`.

-	We then create map of size N to store `node: closest distance from src`.

-	We start a DFS call on source node, keeping track of elapsed time.

-	In the dfs function, we check if elapsed time of node in the map is lower than what is already set (faster path to the node). If not, we break.

-	If yes, we update the time in the map and call dfs for all neighbours of the node.

-	Once all neighbours have been visited, we find the neighbour with max time to reach that node and return it's value.
