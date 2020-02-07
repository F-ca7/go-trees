package main

import (
	"fmt"
	"go-trees/src/avl"
	"go-trees/src/binsearchtree"
	"go-trees/src/btree"
	"go-trees/src/heap"
	"go-trees/src/segment"
	"go-trees/src/trie"
)

func binSearchTest()  {
	arr := []int{9,2,3,5,1,55,100,20,30,910,31,90}

	tree := binsearchtree.BuildSearchTree(arr)
	tree.InOrder()

	var node *binsearchtree.TreeNode
	for i := 0; i < 10; i++ {
		node = tree.Find(i)
		if node == nil {
			fmt.Println(i," Not found")
		} else {
			fmt.Println(i," is found")
		}
	}
	var val int = 9
	if tree.Delete(val) {
		fmt.Printf("删除%d成功\n", val)
	} else {
		fmt.Printf("删除%d失败\n", val)
	}
	tree.InOrder()
}

func heapTest()  {
	arr := []int{9,2,3,5,1,55,100,20,30,910,31,90}
	heap := heap.BuildHeap(arr)
	heap.Push(67)
	heap.Push(999)
	for ; !heap.IsEmpty();  {
		fmt.Printf("%d ", heap.Pop())
	}
}

func segmentTest()  {
	arr := []int{1,3,5,7,9,11}
	tree := segment.BuildSegmentTree(arr)
	tree.Print()
	lb ,rb := 0, 3
	res, err := tree.Query(lb, rb)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d 到 %d 区间的和为 %d", lb, rb, res)
}

func avlTest() {
	arr := []int{9,2,3,5,1,55,100,20,30,910,31,90}
	tree := avl.BuildAvlTree(arr)
	tree.InOrder()
}

func trieTest() {
	words := []string{"app", "apple", "ambulance", "bad", "banana"}
	trieTree := trie.BuildTrieTree(words)
	word := "bad"
	if trieTree.Contains(word) {
		fmt.Println("exist!")
	} else {
		fmt.Println("not exist!")
	}
}

func bTreeTest()  {
	bTree, _ := btree.NewBTree(5)
	_ = bTree.Insert(39, "apple")
	bTree.Insert(22, "bad")
	bTree.Insert(97, "cat")
	bTree.Insert(43, "doom")
	bTree.Insert(53, "ember")
	bTree.Insert(13, "fiend")
	bTree.Insert(21, "grape")
	bTree.Insert(40, "hero")
	fmt.Println("Current key num: ", bTree.Size)
	data, ok := bTree.Get(22)
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(data)
	}
	bTree.Print()

	bTree.DeleteKey(22)
	fmt.Println("删除22后")
	fmt.Println("Current key num: ", bTree.Size)
	data, ok = bTree.Get(22)
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println("found", data)
	}
	bTree.Print()
}

func main() {
	//binSearchTest()
	//heapTest()
	//segmentTest()
	//avlTest()
	//trieTest()
	bTreeTest()

}
