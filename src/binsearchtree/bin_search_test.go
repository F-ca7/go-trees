package main

import (
	"math/rand"
	"testing"
	"time"
)

const size int = 100000

// 查找性能测试
func BenchmarkBinSearchTree_Find(b *testing.B) {
	rand.Seed(time.Now().Unix())
	// 初始化测试数组
	var arr [size]int
	for i := 0 ; i < size; i++ {
		arr[i] = rand.Intn(1<<20)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildSearchTree(arr[:])
	}
}
