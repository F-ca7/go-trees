package heap

// 大根堆
// 完全二叉树 底层使用数组实现
type MaxHeap struct {
	arr []int
}

// 构造大根堆
func BuildHeap(arr []int) *MaxHeap{
	heap := new(MaxHeap)
	heap.arr = arr
	length := len(arr)
	// 自下而上调整堆
	for i := getParentIndex(length-1); i>=0; i--  {
		heap.siftUp(i)
	}
	return heap
}

// 插入元素
func (heap *MaxHeap) Push(val int)  {
	heap.arr = append(heap.arr, val)
	// 自下而上调整堆
	for i := getParentIndex(len(heap.arr)-1); i>=0; i--  {
		heap.siftUp(i)
	}
}

// 弹出根元素
func (heap *MaxHeap) Pop() int{
	root := heap.arr[0]
	length := len(heap.arr)
	heap.arr[0] = heap.arr[length-1]
	heap.siftUp(0)
	heap.arr = heap.arr[:length-1]
	return root
}

func (heap *MaxHeap) IsEmpty() bool{
	return len(heap.arr) == 0
}

// 向上调整
// 使得idx处是最大结点
func (heap *MaxHeap) siftUp(idx int) {
	arr := heap.arr
	var left, right int
	for ; !heap.isLeaf(idx);  {
		left = getLeftIndex(idx)
		right = getRightIndex(idx)

		if right <= len(arr)-1 {
			// 有左右孩子
			if arr[left] >= arr[right] {
				if arr[idx] >= arr[left] {
					// 已经是最大的
					return
				} else {
					swap(arr, idx, left)
					idx = left
				}
			} else {
				if arr[idx] >= arr[right] {
					// 已经是最大的
					return
				} else {
					swap(arr, idx, right)
					idx = right
				}
			}
		} else {
			// 只有左孩子
			if arr[idx] < arr[left] {
				swap(heap.arr, idx, left)
			}
			return
		}
	}
}

func swap(arr []int, i, j int)  {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// 获取父结点索引
func getParentIndex(idx int) int{
	return (idx - 1)/2
}

// 获取左子结点索引
func getLeftIndex(idx int) int{
	return idx*2 + 1
}

// 获取右子结点索引
func getRightIndex(idx int) int{
	return idx*2 + 2
}

// 是否为叶子结点
func (heap MaxHeap) isLeaf(idx int) bool{
	size := len(heap.arr)
	return (idx < size) && (idx >= size/2)
}



