Code
====

```go
func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 1 {
		return []int{0}
	}

	graph := make(map[int][]int, 0)
	order := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
		order[i] = 0
	}

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		order[p[0]] += 1
	}

	queue := []int{}
	solution := []int{}

	for course, order := range order {
		if order == 0 {
			queue = append(queue, course)
			solution = append(solution, course)
		}
	}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		for _, d := range graph[course] {
			order[d] -= 1
			if order[d] == 0 {
				queue = append(queue, d)
				solution = append(solution, d)
			}
		}
	}

	if len(solution) == 0 {
		return solution
	}

	if len(solution) == numCourses {
		return solution
	} else {
		return []int{}
	}
}
```

Solution in mind
================

-	Solution can be obtained by topological sort of DAG.

-	First create graph using the prerequisites and an array called order, which stored number of incoming nodes for a node. The node is stored as index and order as value.

-	The graph is stored as node: all the nodes it points to.

-	Following this, create a queue and place nodes with order 0 in the queue.

-	Loop till the queue is not empty, pop an element and decrease the order for all the elements it points to. If any of these nodes have an order of 0 after decrementing, append them to both solution and the queue.

-	Once the queue is empty, if the graph did not contain a cycle, all nodes will be placed in solution in a topologically sorted order.

-	Verify that the length of this solution is equal to number of nodes, if so solution is valid. If not, no solution exists and we may return an empty list.
