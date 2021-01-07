package topk

import "container/heap"

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
