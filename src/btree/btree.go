package btree

import (
	"errors"
	"fmt"
	"math"
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

	maxKey, minKey := order-1, int(math.Ceil(float64(order)/2)  - 1)
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

	tree.Size++
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

// 从 index+1 处的子结点
// 借一个 kvPair
// 到父结点处
// 父结点再给一个 kvPair
// 到 index 处的子结点
func (parent *TreeNode) borrowFromRight(index int) (err error) {
	if index >= len(parent.children)-1 {
		err = errors.New("no right sibling")
		return
	}
	rightChildNode := parent.children[index+1]

	// 右结点最小的子结点
	toMoveChildNode := rightChildNode.children[0]
	// 父结点给左子结点的键值对
	parentMoveKVPair := parent.kvPairs[index]
	// 要接收的结点
	receiveNode := parent.children[index]
	// 接收父结点来的键值对
	receiveNode.kvPairs[receiveNode.keyNum] = parentMoveKVPair
	receiveNode.keyNum++
	// 右结点最小的子结点给左结点新加进来的kv处
	receiveNode.children[receiveNode.keyNum] = toMoveChildNode
	// 用右节点借出的键值对给父结点
	parent.kvPairs[index] = rightChildNode.kvPairs[0]
	// 从右结点删除掉借出的键值对
	// 同时从右结点删掉给出的最小子结点
	for i := 1; i < rightChildNode.keyNum; i++ {
		rightChildNode.kvPairs[i-1] = rightChildNode.kvPairs[i]
		rightChildNode.children[i-1] = rightChildNode.children[i]
	}
	rightChildNode.children[rightChildNode.keyNum-1] = rightChildNode.children[rightChildNode.keyNum]
	rightChildNode.keyNum--
	rightChildNode.kvPairs[rightChildNode.keyNum] = nil
	rightChildNode.children[rightChildNode.keyNum+1] = nil

	return
}

// 从 index-1 处的子结点
// 借一个 kvPair
// 到父结点处
// 父结点再给一个 kvPair
// 到 index 处的子结点
func (parent *TreeNode) borrowFromLeft(index int) (err error) {
	if index <= 0 || parent.keyNum < 2 {
		err = errors.New("no left sibling")
		return
	}
	leftChildNode := parent.children[index-1]

	// 左结点最大的子结点
	toMoveChildNode := leftChildNode.children[leftChildNode.keyNum]
	// 父结点给右子结点的键值对
	parentMoveKVPair := parent.kvPairs[index-1]
	// 要接收的结点
	receiveNode := parent.children[index]
	// 接收父结点来的键值对 以及 左结点的最大子结点
	// 先整体右移一格
	receiveNode.children[receiveNode.keyNum+1] = receiveNode.children[receiveNode.keyNum]
	for i := receiveNode.keyNum; i > 0; i-- {
		receiveNode.kvPairs[i] = receiveNode.kvPairs[i-1]
		receiveNode.children[i] = receiveNode.children[i-1]
	}
	// 再接收
	receiveNode.kvPairs[0] = parentMoveKVPair
	receiveNode.children[0] = toMoveChildNode
	receiveNode.keyNum++
	// 右结点最小的子结点给左结点新加进来的kv处
	receiveNode.children[receiveNode.keyNum] = toMoveChildNode
	// 用左节点借出的键值对给父结点
	parent.kvPairs[index-1] = leftChildNode.kvPairs[leftChildNode.keyNum-1]
	// 从左结点删除掉借出的键值对
	// 同时从左结点删掉给出的最小子结点
	leftChildNode.keyNum--
	leftChildNode.kvPairs[leftChildNode.keyNum] = nil
	leftChildNode.children[leftChildNode.keyNum+1] = nil

	return
}

// 合并父结点下
// index 和 index-1 处的结点
// 把右子结点合并给左子结点
func (parent *TreeNode) leftMerge(index int) (err error) {
	if index > parent.keyNum || index <= 0{
		err = errors.New("key index out of bound")
		return
	}
	leftChildNode := parent.children[index-1]
	rightChildNode := parent.children[index]
	// 父结点给左结点的键值对
	toMoveKVPair := parent.kvPairs[index-1]
	leftChildNode.kvPairs[leftChildNode.keyNum] = toMoveKVPair
	leftChildNode.keyNum++
	// 更新父结点
	for i := index; i < parent.keyNum; i++  {
		parent.kvPairs[i-1] = parent.kvPairs[i]
		parent.children[i] = parent.children[i + 1]
	}
	parent.kvPairs[parent.keyNum-1] = nil
	parent.children[parent.keyNum] = nil
	parent.keyNum--
	// 把右结点键值对合并到左结点
	for i := 0; i < rightChildNode.keyNum; i++ {
		leftChildNode.kvPairs[leftChildNode.keyNum + i] = rightChildNode.kvPairs[i]
	}
	// 把右结点的子结点合并到左结点
	for i := 0; i < rightChildNode.keyNum + 1; i++ {
		leftChildNode.children[leftChildNode.keyNum + i] = rightChildNode.children[i]
	}
	leftChildNode.keyNum += rightChildNode.keyNum
	return
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

// 删除 key
// 返回 key 是否存在
func (tree *BTree) DeleteKey(key int) bool {
	if tree.root == nil {
		return false
	}
	tree.deleteHelper(key, tree.root)
	return true
}

func (tree *BTree) deleteHelper(key int, node *TreeNode) *TreeNode {
	toDeleteIdx := node.binarySearchPos(key)

	if node.isLeaf {
		// 如果是叶子结点
		// 找到则直接删除
		if toDeleteIdx < node.keyNum && key == node.kvPairs[toDeleteIdx].key {
			for i := toDeleteIdx; i < node.keyNum - 1; i++ {
				node.kvPairs[i] = node.kvPairs[i+1]
			}
			node.kvPairs[node.keyNum-1] = nil
			node.keyNum--
			tree.Size--
		}
		return node
	}
	// 如果不是叶子结点
	if toDeleteIdx < node.keyNum && key == node.kvPairs[toDeleteIdx].key {
		// 找到了 key
		// 用 删除处的左链接子结点的最右键值对 来替换
		// 即前继键值对
		leftChildNode := node.children[toDeleteIdx]
		predecessorKVPair := leftChildNode.kvPairs[leftChildNode.keyNum-1]
		// 删除前继结点并记录
		node = tree.deleteHelper(predecessorKVPair.key, node)

		node.kvPairs[toDeleteIdx] = predecessorKVPair
	} else {
		// 递归到子结点去删除
		keyDeletedNode := tree.deleteHelper(key, node.children[toDeleteIdx])
		// 调整状态
		if keyDeletedNode.keyNum < tree.minKeyNum {
			// 发生下溢出
			if node.childHasLeftSiblingAt(toDeleteIdx) {
				// 优先考虑左兄弟结点
				if node.children[toDeleteIdx-1].keyNum > tree.minKeyNum {
					// 左兄弟的键值对足够
					// 直接借出
				 	_ = node.borrowFromLeft(toDeleteIdx)
				} else {
					// 左兄弟的键值对不够
					// 左合并
					_ = node.leftMerge(toDeleteIdx)
				}
			} else {
				if node.children[toDeleteIdx].keyNum > tree.minKeyNum {
					// 右兄弟的键值对足够
					// 直接借出
					_ = node.borrowFromRight(toDeleteIdx)
				} else {
					// 右兄弟的键值对不够
					// 右合并
					_ = node.leftMerge(toDeleteIdx + 1)
				}
			}
		}
	}
	return node
}

// parent 的 index 处子结点
// 是否有左兄弟结点
func (parent *TreeNode) childHasLeftSiblingAt(index int) bool {
	return index > 0 && index <= parent.keyNum
}

// parent 的 index 处子结点
// 是否有右兄弟结点
func (parent *TreeNode) childHasRightSiblingAt(index int) bool {
	return index >= 0 && index < parent.keyNum
}

// 根据 key 获取值
// ok 表示 key 是否存在
func (tree *BTree) Get(key int) (value string, ok bool) {
	current := tree.root
	var index int
	for ; current != nil;  {
		index = current.binarySearchPos(key)
		// 判断索引处是否匹配
		if index < current.keyNum && key == current.kvPairs[index].key {
			value = current.kvPairs[index].value
			ok = true
			return
		} else {
			// 不匹配则往子结点查找
			current = current.children[index]
		}
	}
	ok = false
	return
}



// 打印B- 树
// 调试用
func (tree *BTree) Print() {
	if tree.Size == 0 {
		fmt.Println("Empty tree")
	} else {
		printHelper(tree.root, 1, 0, 0)
	}

}

func printHelper(node *TreeNode, level, pIdx, cIdx int) {
	currentNode := node
	if currentNode == nil {
		return
	}
	fmt.Println("Current level: ", level)
	fmt.Println("Parent index: ", pIdx)
	fmt.Println("Child index: ", cIdx)

	fmt.Print("Current keys: ")
	for i := 0; i < len(currentNode.kvPairs); i++ {
		fmt.Print(currentNode.kvPairs[i], " ")
	}
	fmt.Println("\n-----------------")
	if !currentNode.isLeaf {
		for i, v := range currentNode.children {
			printHelper(v, level+1, i, cIdx)
		}
	}
}
