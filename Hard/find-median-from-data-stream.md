Code
====

```go
type MedianFinder struct {
	arr []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{}
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, num)
	} else {
		i := sort.Search(len(this.arr), func(i int) bool { return num < this.arr[i] })
		rArr := append([]int{}, this.arr[i:]...)
		this.arr = append(this.arr[:i], num)
		this.arr = append(this.arr, rArr...)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.arr)
	if n%2 == 0 {
		return float64(this.arr[n/2]+this.arr[(n/2)-1]) / 2
	} else {
		return float64(this.arr[n/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
```

Solution in mind
================

-	Insert into slice using binary search
-	Compute median using it's formula
