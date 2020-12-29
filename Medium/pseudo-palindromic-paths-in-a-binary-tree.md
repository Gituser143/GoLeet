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

type data struct {
	nodes  map[int]int
	numOdd int
}

func traverse(root *TreeNode, d0 data) int {
	if root == nil {
		return 0
	}

	d := data{nodes: make(map[int]int, 0), numOdd: d0.numOdd}

	for key, val := range d0.nodes {
		d.nodes[key] = val
	}

	if val, ok := d.nodes[root.Val]; ok {
		d.nodes[root.Val] += 1
		if (val+1)%2 == 0 {
			d.numOdd -= 1
		} else {
			d.numOdd += 1
		}
	} else {
		d.nodes[root.Val] = 1
		d.numOdd += 1
	}

	if root.Right == nil && root.Left == nil {
		if d.numOdd <= 1 {
			return 1
		} else {
			return 0
		}
	}

	l := traverse(root.Left, d)
	r := traverse(root.Right, d)

	return l + r
}

func pseudoPalindromicPaths(root *TreeNode) int {
	d := data{nodes: make(map[int]int, 0), numOdd: 0}

	num := traverse(root, d)

	return num
}
```

Solution in mind
================

-	Iterate through nodes of the tree using a depth first strategy.

-	Track values encountered on each path and maintain a count of odd count values.

-	Here, the number of times a value is encountered is maintained by `data.nodes`. `nodes`, is a map of the format `Node.Val: count`. Additionally, an integer is maintained to keep track of number of odd occurrences (`data.numOdd`).

-	On reaching a leaf node, if the number of odd occurrences is lesser than or equal to 1, we have a Pseudo Palindromic Path, thus we return 1, else return 0.
