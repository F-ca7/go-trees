package avl

import (
	"math/rand"
	"testing"
	"time"
)

const size int = 100000

func BenchmarkBuildAvlTree(b *testing.B) {
	rand.Seed(time.Now().Unix())
	// 初始化测试数组
	var arr []int = make([]int, size)
	for i := 0 ; i < size; i++ {
		arr[i] = rand.Intn(1<<20)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BuildAvlTree(arr)
	}
}
