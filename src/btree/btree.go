package btree

import (
	"errors"
	"fmt"
)

type KVPair struct {
	key int
	value string
}

func (x *KVPair) CompareTo(y *KVPair) int {
	if x.key == y.key {
		return 0
	} else if x.key > y.key {
		return 1
	} else {
		return -1
	}
}

type TreeNode struct {
	kvPairs  []*KVPair    // 键值对
	children []*TreeNode // 子结点
	isLeaf   bool        // 是否为叶子结点
	keyNum   int         // 该结点的 key 数量
}

type BTree struct {
	root *TreeNode
	order int // 阶数
	maxKeyNum int // key 达到上限就需要分裂
	minKeyNum int // 结点 key 数的下限
	Size int // 总 key 数目
}

func (tree *BTree) NewNode() *TreeNode {
	keys := make([]*KVPair, tree.maxKeyNum)
	children := make([]*TreeNode, tree.order)
	return &TreeNode{
		kvPairs:  keys,
		children: children,
		isLeaf:   true,
		keyNum:   0,
	}
}

func NewBTree(order int) (tree *BTree, err error) {
	if order < 3 {
		err = errors.New("order no less than 3")
		return
	}
	maxKey, minKey := order-1, order/2 - 1
	tree = &BTree{order: order, maxKeyNum:maxKey, minKeyNum:minKey}
	// 初始化根节点
	tree.root = tree.NewNode()
	return
}

// 向B树插入数据
func (tree *BTree) Insert(key int, data string) (err error) {
	tree.root = tree.insertHelper(key, data, tree.root)
	return
}

func (tree *BTree) insertHelper(key int, data string, node *TreeNode) *TreeNode {
 	if node.keyNum == tree.maxKeyNum {
		// 结点的key已满
		// 分裂
		node = tree.split(node)
	}
	toInsertIndex := node.binarySearchPos(key)
	kvPair := node.kvPairs[toInsertIndex]

	if kvPair != nil && kvPair.key == key {
		// 如果该 key 已存在
		// 直接替换 value
		kvPair.value = data
		return node
	}

	if node.isLeaf {
		// 如果是叶子结点
		// 直接插入
		insertKV := &KVPair{key:key, value:data}
		for i := node.keyNum; i > toInsertIndex; i-- {
			node.kvPairs[i] = node.kvPairs[i-1]
		}
		node.kvPairs[toInsertIndex] = insertKV
		node.keyNum++
		return node
	}

	insertedNode := tree.insertHelper(key, data, node.children[toInsertIndex])
	if insertedNode.keyNum == 1 {
		// 返回的结点发生了分裂
		// 需要合并到当前结点
		for i := node.keyNum; i > toInsertIndex; i-- {
			node.kvPairs[i] = node.kvPairs[i-1]
		}
		node.kvPairs[toInsertIndex] = insertedNode.kvPairs[0]
		// 合并插入结点的两个子结点
		for i := node.keyNum + 1; i > toInsertIndex + 1; i-- {
			node.children[i] = node.children[i-1]
		}
		node.children[toInsertIndex] = insertedNode.children[0]
		node.children[toInsertIndex + 1] = insertedNode.children[1]

		node.keyNum++
	}
	return node
}

// 将结点 node 分裂为
// 一个父结点 parent 和 两个子结点
// 返回 parent
func (tree *BTree) split(node *TreeNode) (parent *TreeNode) {
	// 满 key 数
	maxKeyNum := tree.maxKeyNum
	parent = tree.NewNode()
	leftChild := tree.NewNode()
	rightChild := tree.NewNode()
	// 原来是叶子
	// 则现在也是叶子
	leftChild.isLeaf = node.isLeaf
	rightChild.isLeaf = node.isLeaf

	mid := maxKeyNum/2
	leftChild.keyNum = mid
	rightChild.keyNum = maxKeyNum - mid - 1
	// 原结点的左半部分给左孩子
	for i := 0; i < mid; i++ {

		leftChild.kvPairs[i] = node.kvPairs[i]
		leftChild.children[i] = node.children[i]
	}
	leftChild.children[mid] = node.children[mid]
	// 原结点的右半部分给左孩子
	for i := mid + 1; i < maxKeyNum; i++ {
		rightChild.kvPairs[i-mid-1] = node.kvPairs[i]
		rightChild.children[i-mid-1] = node.children[i]
	}
	rightChild.children[maxKeyNum- mid-1] = node.children[maxKeyNum]
	// mid 处变为父结点
	parent.kvPairs[0] = node.kvPairs[mid]
	parent.isLeaf = false
	parent.keyNum = 1
	parent.children[0] = leftChild
	parent.children[1] = rightChild
	return
}

// 通过二分查找找到
// key 在 node 中的位置
// 或 key 插入 node 的位置
func (node *TreeNode) binarySearchPos(key int) int {
	kvPairs := node.kvPairs
	startPos, endPos := 0, node.keyNum - 1
	mid := (startPos + endPos)/2
	var midKey int
	for ; startPos <= endPos;  {
		midKey = kvPairs[mid].key
		if key == midKey {
			return mid
		} else if key < midKey {
			endPos = mid - 1
		} else {
			startPos = mid + 1
		}
		mid = (startPos + endPos)/2

	}
	return startPos
}

// 根据 key 获取值
// 若不存在 则 err
func (tree *BTree) Get(key int) (value string, err error) {
	current := tree.root
	var index int
	for ; current != nil;  {
		index = current.binarySearchPos(key)
		// 判断索引处是否匹配
		if index < current.keyNum && key == current.kvPairs[index].key {
			value = current.kvPairs[index].value
			return
		} else {
			// 不匹配则往子结点查找
			current = current.children[index]
		}
	}
	err = errors.New(fmt.Sprintf("key %d not found", key))
	return
}

// 打印B- 树
// 调试用
func (tree *BTree) Print() {

}
