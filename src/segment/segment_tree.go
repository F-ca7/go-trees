package segment

import (
	"errors"
	"fmt"
	"math"
)

// 以 结点存储信息为区间和 为例
type SegmentTree struct {
	arr []int
	height int	// 树的高度
	length int // 原数组长度
}

func BuildSegmentTree(src []int)  *SegmentTree{
	tree := new(SegmentTree)
	length := len(src)
	tree.height = int(math.Ceil(math.Log2(float64(length)))) + 1
	tree.arr = make([]int, 1<<uint(tree.height) - 1)
	tree.length = len(src)

	tree.buildTreeHelper(src, 0, 0, length-1)
	return tree
}

// 查询子区间的和
// 从 leftBound 到 rightBound
func (tree *SegmentTree) Query(leftBound, rightBound int) (sum int, err error){
	sum = 0
	if leftBound > rightBound {
		err = errors.New("leftBound is greater than rightBound")
		return
	}
	if rightBound > tree.length-1 || leftBound < 0 {
		err = errors.New("out of bound")
		return
	}

	sum = tree.queryHelper(0, 0, tree.length-1, leftBound, rightBound)
	return
}


func (tree *SegmentTree) queryHelper(node, start ,end, leftBound, rightBound int) int {
	if leftBound <= start && rightBound >= end {
		// 查询区间包括了子结点区间
		// 直接返回
		return tree.arr[node]
	}
	mid	:= (start + end)/2
	left := getLeftIndex(node)
	right := getRightIndex(node)
	var leftSum, rightSum int
	if leftBound <= mid {
		// 要往左子树找
		leftSum = tree.queryHelper(left, start, mid, leftBound, rightBound)
	}
	if rightBound >= (mid + 1) {
		// 要往右子树找
		rightSum = tree.queryHelper(right, mid + 1, end, leftBound, rightBound)
	}
	return leftSum + rightSum;
}

// 更新原数组指定位置的值
func (tree *SegmentTree) UpdateAt(idx, val int)  {
	tree.updateHelper(0, 0, tree.length-1, idx, val)
}

func (tree *SegmentTree) updateHelper(node, start, end, idx, val int) {
	if start == end {
		// 更新结点
		tree.arr[node] = val
		return
	}
	mid := (start + end)/2
	left := getLeftIndex(node)
	right := getRightIndex(node)
	if idx >= start && idx <= mid {
		// 更新结点在左子树
		tree.updateHelper(left, start, mid, idx, val)
	} else {
		// 更新结点在右子树
		tree.updateHelper(right, mid + 1, end, idx, val)
	}
	tree.arr[node] = tree.arr[left] + tree.arr[right]
}

// 递归构造线段树
func (tree *SegmentTree) buildTreeHelper(src []int, node, start, end int)  {
	if start == end {
		tree.arr[node] = src[start]
		return
	}
	mid := (start + end)/2
	left := getLeftIndex(node)
	right := getRightIndex(node)

	tree.buildTreeHelper(src, left, start, mid)
	tree.buildTreeHelper(src, right, mid + 1, end)
	tree.arr[node] = tree.arr[left] + tree.arr[right]
}

func (tree *SegmentTree) Print() {
	fmt.Println(tree.arr)
}

// 获取左子结点索引
func getLeftIndex(idx int) int{
	return idx*2 + 1
}

// 获取右子结点索引
func getRightIndex(idx int) int{
	return idx*2 + 2
}