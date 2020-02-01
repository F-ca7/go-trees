package avl

import "fmt"

type AvlNode struct {
	Value int
	height int // 结点高度
	left *AvlNode
	right *AvlNode
}

type AvlTree struct {
	root *AvlNode
}

// 构造平衡二叉树
func BuildAvlTree(src []int) *AvlTree {
	tree := new(AvlTree)
	for _, v := range src{
		tree.Insert(v)
	}
	return tree
}

// 插入结点
func (tree *AvlTree) Insert(val int)  {
	newNode := &AvlNode{Value: val, height: 1}
	if tree.root == nil {
		tree.root = newNode
	} else {
		// 找到插入的位置
		current := tree.root
		for {
			if val < current.Value {
				if current.left == nil {
					current.left = newNode
					break
				}
				current = current.left
			} else {
				if current.right == nil {
					current.right = newNode
					break
				}
				current = current.right
			}
		}
		// 插入完后开始调节
		newNode.adjust()
	}
}

// 以 axisNode 为轴进行左旋
// 返回 subRoot 为旋转后的根节结点
func (axisNode *AvlNode) leftRotate() (subRoot *AvlNode) {
	subRoot = axisNode.right
	axisNode.right = subRoot.left
	subRoot.left = axisNode
	// 从下往上更新高度
	axisNode.updateHeight()
	subRoot.updateHeight()
	return
}

// 以 axisNode 为轴进行右旋
// 返回 subRoot 为旋转后的根节结点
func (axisNode *AvlNode) rightRotate() (subRoot *AvlNode) {
	subRoot = axisNode.left
	axisNode.left = subRoot.right
	subRoot.right = axisNode
	// 从下往上更新高度
	axisNode.updateHeight()
	subRoot.updateHeight()
	return
}

// 先左旋 再右旋
// 返回 subRoot 为旋转后的根节结点
func (axisNode *AvlNode) leftRightRotate() (subRoot *AvlNode) {
	// 对 axis 的左结点先左旋转
	newAxisLeft := axisNode.left.leftRotate()
	axisNode.left = newAxisLeft
	// 再绕 axis 右旋转
	subRoot = axisNode.rightRotate()
	return
}

// 先右旋 再左旋
// 返回 subRoot 为旋转后的根节结点
func (axisNode *AvlNode) rightLeftRotate() (subRoot *AvlNode) {
	// 对 axis 的右结点先右旋转
	newAxisRight := axisNode.right.rightRotate()
	axisNode.right = newAxisRight
	// 再绕 axis 左旋转
	subRoot = axisNode.leftRotate()
	return
}

// 根据平衡情况来旋转调整
func (node *AvlNode) adjust() {
	rlDiff := node.right.getHeight() - node.left.getHeight()
	if rlDiff == 2 {
		if node.right.right.getHeight() > node.right.left.getHeight() {
			node.leftRotate()
		} else {
			node.rightLeftRotate()
		}
	} else if rlDiff == -2 {
		if node.left.left.getHeight() > node.left.right.getHeight() {
			node.rightRotate()
		} else {
			node.rightLeftRotate()
		}
	}
}

// 更新 node 的高度
func (node *AvlNode) updateHeight()  {
	node.height = max(node.left.Value, node.right.height) + 1
}

// 获取 node 的高度
func (node *AvlNode) getHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (tree *AvlTree) InOrder() {
	if tree.root != nil {
		tree.root.InOrder()
	}
}

func (node *AvlNode) InOrder() {
	if node.left!=nil {
		node.left.InOrder()
	}
	fmt.Printf("%d ", node.Value)
	if node.right!=nil {
		node.right.InOrder()
	}
}
