package main

import "fmt"

func main() {
	arr := [...]int{9,2,3,5,1,55,100,20,30,910,31,90}

	tree := BuildSearchTree(arr[:])
	tree.InOrder()

	//var node *TreeNode
	//for i := 0; i < 10; i++ {
	//	node = tree.Find(i)
	//	if node == nil {
	//		fmt.Println(i," Not found")
	//	} else {
	//		fmt.Println(i," is found")
	//	}
	//}
	var val int = 9
	if tree.Delete(val) {
		fmt.Printf("删除%d成功\n", val)
	} else {
		fmt.Printf("删除%d失败\n", val)
	}
	tree.InOrder()
}