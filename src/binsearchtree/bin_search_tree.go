package main

import "fmt"

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

// 先序遍历
func (node *TreeNode) PreOrder() {
	fmt.Printf("%d ", node.Value)
	if node.Left!=nil {
		node.Left.PreOrder()
	}
	if node.Right!=nil {
		node.Right.PreOrder()
	}
}

// 后序遍历
func (node *TreeNode) PostOrder() {
	if node.Left!=nil {
		node.Left.PostOrder()
	}
	if node.Right!=nil {
		node.Right.PostOrder()
	}
	fmt.Printf("%d ", node.Value)
}

// 中序遍历
func (node *TreeNode) InOrder() {
	if node.Left!=nil {
		node.Left.InOrder()
	}
	fmt.Printf("%d ", node.Value)
	if node.Right!=nil {
		node.Right.InOrder()
	}

}

type BinSearchTree struct {
	root *TreeNode
}

func (tree *BinSearchTree) Insert(value int)  {
	newNode := &TreeNode{value, nil, nil}
	if tree.root==nil {
		tree.root = newNode
	} else {
		// 找到插入的位置
		current := tree.root
		for {
			if value < current.Value {
				if current.Left==nil {
					current.Left = newNode
					return
				}
				current = current.Left
			} else {
				if current.Right==nil {
					current.Right = newNode
					return
				}
				current = current.Right
			}
		}
	}

}


// 构造二叉搜索树
func BuildSearchTree(arr []int) *BinSearchTree{
	fmt.Println(arr)
	root := new(BinSearchTree)
	for _,v := range arr {
		root.Insert(v)
	}
	return root
}

// 先序遍历
func (tree *BinSearchTree) PreOrder()  {
	if tree.root != nil {
		tree.root.PreOrder()
	}
	fmt.Println()
}

// 后序遍历
func (tree *BinSearchTree) PostOrder()  {
	if tree.root != nil {
		tree.root.PostOrder()
	}
	fmt.Println()
}

// 中序遍历，从小到大的顺序
func  (tree *BinSearchTree) InOrder()  {
	if tree.root != nil {
		tree.root.InOrder()
	}
	fmt.Println()
}

// 查找指定值的第一个结点
func (tree *BinSearchTree) Find(val int)  *TreeNode{
	node := tree.root
	for ; node != nil && node.Value != val; {
		if val < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

// 删除指定值的第一个结点
func (tree *BinSearchTree) Delete(val int)  bool{
	node := tree.root
	var parent *TreeNode
	// 是否为父结点的左孩子
	isLeftChild := false
	// 先找到要删除的结点及其父结点
	for ; node != nil && node.Value != val; {
		parent = node
		if val < node.Value {
			node = node.Left
			isLeftChild = true
		} else if val > node.Value {
			node = node.Right
			isLeftChild = false
		}
	}
	if node == nil {
		return false
	}
	// 找到了要删除的结点node
	if node.Left == nil && node.Right == nil {
		// 叶子结点
		// 直接删除
		if isLeftChild {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return true
	}
	if node.Left == nil && node.Right != nil {
		// 只有一个右孩子
		// 直接给parent
		if isLeftChild {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		node = nil
		return true
	}
	if node.Left != nil && node.Right == nil {
		// 只有一个左孩子
		// 直接给parent
		if isLeftChild {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
		node = nil
		return true
	}
	// node有两个孩子
	// 找到中序遍历的后续结点与node交换
	inOrderSuccessor, successorParent, isSuccessorRight := node.getInOrderSuccessor()
	// 直接交换值
	node.Value = inOrderSuccessor.Value
	// Successor没有孩子或者只有右孩子
	if inOrderSuccessor.Right != nil {
		// 只有右孩子
		if isSuccessorRight {
			successorParent.Right = inOrderSuccessor.Right
		} else {
			successorParent.Left = inOrderSuccessor.Right
		}
	} else {
		// 没有孩子
		if isSuccessorRight {
			successorParent.Right = nil
		} else {
			successorParent.Left = nil
		}
	}
	inOrderSuccessor = nil
	return true
}


// 获取结点中序遍历的后续结点
// 该结点应该有左右孩子
func (node *TreeNode)getInOrderSuccessor() (successor, parent *TreeNode, isSuccessorRight bool){
	isSuccessorRight = true
	parent = node
	// 右孩子子树的最左结点
	successor = node.Right
	for ; successor.Left != nil;  {
		isSuccessorRight = false
		parent = successor
		successor = successor.Left
	}
	return
}