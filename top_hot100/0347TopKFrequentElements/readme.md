# 347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:
```
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
```
示例 2:
```
输入: nums = [1], k = 1
输出: [1]
```

## 解法

```go
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		numMap[num] = numMap[num] + 1
	}

	h := &Iheap{}
	heap.Init(h)
	for key, value := range numMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).([2]int)[0])
	}

	return result
}

// Iheap 堆
type Iheap [][2]int

func (h Iheap) Len() int           { return len(h) }
func (h Iheap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h Iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push push
func (h *Iheap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop pop
func (h *Iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
```