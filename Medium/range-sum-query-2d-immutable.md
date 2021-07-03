Code
====

```go
type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := NumMatrix{Matrix: matrix}
	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.Matrix[i][j]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
```
