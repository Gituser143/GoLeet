Code
====

```go

	func checkCyclicParent(cur int, leftChild []int, rightChild []int, parentsList map[int]bool) bool {
		if _, ok := parentsList[cur]; ok {
			fmt.Println(ok, parentsList[cur])
			return false
		} else {
			if cur != -1 {
				parentsList[cur] = true
				return checkCyclicParent(leftChild[cur], leftChild, rightChild, parentsList) && checkCyclicParent(rightChild[cur], leftChild, rightChild, parentsList)
			}
			return true
		}
	}

	func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {

		// Check for multiple parents
		parents := make(map[int]bool, 0)

		for i := 0; i < n; i++ {
			parents[i] = false
		}

		for i := 0; i < n; i++ {
			if leftChild[i] != -1 {
				if val, ok := parents[leftChild[i]]; ok && val {
					return false
				} else {
					parents[leftChild[i]] = true
				}
			}

			if rightChild[i] != -1 {
				if val, ok := parents[rightChild[i]]; ok && val {
					return false
				} else {
					parents[rightChild[i]] = true
				}
			}
		}

		fmt.Println("Valid single parents")

		// Check if all children have parents
		noParent := false
		for child, parent := range parents {
			fmt.Println(child, parent)
			if !parent {
				if noParent {
					return false
				}
				noParent = true
			}
		}

		fmt.Println("Valid single orphan")
		parentsList := make(map[int]bool, 0)
		return checkCyclicParent(0, leftChild, rightChild, parentsList)
	}
```

Solution in mind
================

-	Ensure children have single parents. Do so my creating a map which stores `child: true/false`. If we ever encounter a value of `child: true`. This means the child was already traversed by a different parent.

-	Only one child can have no parent, if more than one occur, there are 2 root nodes. This can be checked by verifying that no more than 1 key value pair of `child: false` exist in the above map.

-	Check for cycles. If we ever encounter an already parsed parent while traversing the tree, a cycle exists.
