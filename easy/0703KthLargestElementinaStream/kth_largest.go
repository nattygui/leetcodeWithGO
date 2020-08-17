package _703KthLargestElementinaStream

type KthLargest struct {
	size  int //堆的容量，不包括数组第一个元素
	data  []int
	count int //当前元素数量
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	kth.size = k
	kth.data = []int{0}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.count < this.size-1 {
		this.data = append(this.data, val)
		this.count += 1
	} else if this.count == this.size-1 {
		this.data = append(this.data, val)
		this.count += 1
		//第一次填满k个数，使堆有序，后续add只需在data[1]覆盖放入新值
		n := len(this.data) - 1 //data的最大下标
		for i := n / 2; i >= 1; i-- {
			heapify(this.data, i)
		}
	} else {
		//如果新元素大于堆中第k大的元素，说面目前的堆顶一顶不是第k大，直接用新元素替换堆顶
		if val > this.data[1] {
			this.data[1] = val
			heapify(this.data, 1)
		}
	}
	//fmt.Println("add val:",val,this.data)
	return this.data[1]
}

//heapify 从给定的i向下堆化为小顶堆
func heapify(a []int, i int) {
	n := len(a) - 1
	for {
		minPos := i
		if i*2 <= n && a[i*2] < a[minPos] {
			minPos = i * 2
		}
		if i*2+1 <= n && a[i*2+1] < a[minPos] {
			minPos = i*2 + 1
		}
		if minPos == i {
			break
		}
		a[minPos], a[i] = a[i], a[minPos]
		i = minPos
	}
}
