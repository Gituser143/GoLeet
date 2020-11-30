Code
====

```go
type point struct {
	x       int
	height  int
	isStart bool
}

func getSkyline(buildings [][]int) [][]int {
	// make a slice to store all points
	pointList := make([]point, 0)

	// Populate pointList
	for _, building := range buildings {
		start, end, height := building[0], building[1], building[2]
		p1 := point{start, height, true}
		p2 := point{end, height, false}

		pointList = append(pointList, p1, p2)
	}

	// Sort list based on specified conditions
	sort.Slice(pointList, func(i, j int) bool {
		// pick lower of x value, if same follow below rules
		if pointList[i].x > pointList[j].x {
			return false
		} else if pointList[i].x < pointList[j].x {
			return true
		} else {
			// If both are start points, pick one with higher height
			// If both are end, pick one with lower height
			// If both are at same point and either one is start, return the one which is start
			if pointList[i].isStart && pointList[j].isStart {
				if pointList[i].height > pointList[j].height {
					return true
				} else if pointList[i].height < pointList[j].height {
					return false
				} else {
					return pointList[i].isStart
				}
			} else if !pointList[i].isStart && !pointList[j].isStart {
				if pointList[i].height > pointList[j].height {
					return false
				} else if pointList[i].height < pointList[j].height {
					return true
				} else {
					return pointList[i].isStart
				}
			} else {
				return pointList[i].isStart
			}
		}
	})

	// Maintain a map to keep track of points
	heightQueue := make(map[int]int, 0)
	maxInQueue := 0
	heightQueue[0] = 1

	sol := [][]int{}

	// Iterate through points
	for _, point := range pointList {

		// If point is starting of building, add to queue
		// If the point is taller than all existing ones, append to solution and modify height
		if point.isStart {
			if _, ok := heightQueue[point.height]; ok {
				heightQueue[point.height] += 1
			} else {
				heightQueue[point.height] = 1
			}

			if point.height > maxInQueue {
				sol = append(sol, []int{point.x, point.height})
				maxInQueue = point.height
			}
		} else {
			// If point is ending of building, decrement or delte value from map
			val := heightQueue[point.height]

			if val == 1 {
				delete(heightQueue, point.height)
			} else {
				heightQueue[point.height] -= 1
			}

			// If point was the sole maximum, compute new max and append to solution
			if point.height == maxInQueue {
				if _, ok := heightQueue[point.height]; !ok {
					max := 0
					for key, _ := range heightQueue {
						if key > max {
							max = key
						}
					}
					maxInQueue = max
					sol = append(sol, []int{point.x, maxInQueue})
				}
			}

		}
	}

	return sol

}
```

Solution in mind
================

-	Solution based off [YouTube video](https://youtu.be/GSBLe8cKu0s).
