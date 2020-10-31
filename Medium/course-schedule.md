Code
====

```go
func isCyclicUtil(node int, graph map[int][]int, visited, recStack []bool) ([]bool, []bool, bool) {
	visited[node] = true
	recStack[node] = true

	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			visited, recStack, isCycle := isCyclicUtil(neighbour, graph, visited, recStack)
			if isCycle {
				return visited, recStack, true
			}
		} else if recStack[neighbour] {
			return visited, recStack, true
		}
	}

	recStack[node] = false
	return visited, recStack, false
}

func isCyclic(graph map[int][]int) bool {
	visited := make([]bool, len(graph))
	recStack := make([]bool, len(graph))

	isCycle := false
	for node, _ := range graph {
		if !visited[node] {
			visited, recStack, isCycle = isCyclicUtil(node, graph, visited, recStack)
			if isCycle {
				return true
			}
		}
	}

	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 1 {
		return true
	}

	graph := make(map[int][]int, 0)

	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}

	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}

	if isCyclic(graph) {
		return false
	} else {
		return true
	}
}
```

Solution in mind
================

-	The problem boils down to checking if there is a cycle in the graph, if yes, courses can't be finished.

-	To check for cycle, we do so using dfs. We call check cycle for all nodes in graph.

-	For each check cycle call, we maintain a recursion stack and visited.

-	Visted is to keep track of previously visited nodes. The recursion stack is required if two nodes lead to the same node but do not form a cycle.

-	We need a recursion stack because if a neighbour was visited, it need not be from the current path we traverse on. Only if the neighbour was visited and it exists in recursion stack, can we say that a cycle exists.
