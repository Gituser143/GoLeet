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

var min_x, max_x int

func traverse(root *TreeNode, x, y int, levels map[int]map[int][]int) map[int]map[int][]int {
	if root == nil {
		return levels
	}

	if _, ok := levels[x]; ok {
		if _, ok = levels[x][y]; ok {
			levels[x][y] = append(levels[x][y], root.Val)
		} else {
			levels[x][y] = []int{root.Val}
		}
	} else {
		levels[x] = make(map[int][]int, 0)
		levels[x][y] = []int{root.Val}
	}

	if x > max_x {
		max_x = x
	}

	if x < min_x {
		min_x = x
	}

	levels = traverse(root.Left, x-1, y+1, levels)
	levels = traverse(root.Right, x+1, y+1, levels)

	return levels
}

func verticalTraversal(root *TreeNode) [][]int {
	min_x = 0
	max_x = 0
	levels := make(map[int]map[int][]int, 0)

	levels = traverse(root, 0, 0, levels)

	sol := make([][]int, max_x-min_x+1)

	for key, nums := range levels {
		x := key - min_x
		row := []int{}

		keys := []int{}
		for key, _ := range nums {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		for _, key := range keys {
			sort.Ints(nums[key])
			row = append(row, nums[key]...)
		}

		sol[x] = row
	}

	return sol
}
```

Solution in mind
================

-	We maintain a map of format `map[x][y] = [ints]`, which stores values of nodes in their corresponding position (x, y).

-	We index the root as `0, 0`, a left child as `x-1, y+1` and a right child as `x+1, y+1`.

-	We traverse the tree and populate the map. During population, we keep track of the minimum and maximum values `x` takes. This will later help us calculate the total number of rows the solution has.

-	We iterate through each `map[x]`, the values in all of it's `y` combined will give us the row at index `x - min_x`.

-	But to populate this row, we must iterate through `y` in a sorted order. We must additionally ensure that `map[x][y]` when added to the row, is sorted in itself.
