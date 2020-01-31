package segment

import (
	"math/rand"
	"testing"
	"time"
)

const (
	SIZE int = 10000
	TEST_COUNT int = 100
)

// 线段树查询功能测试
func TestSegmentTree_Query(t *testing.T) {
	rand.Seed(time.Now().Unix())
	// 初始化测试数组
	var arr [SIZE]int
	for i := 0 ; i < SIZE; i++ {
		arr[i] = rand.Intn(1<<10)
	}
	tree := BuildSegmentTree(arr[:])
	for i := 0; i < TEST_COUNT; i++  {
		lb := rand.Intn(SIZE-1)
		rb := rand.Intn(SIZE - lb) + lb
		sum1, err := tree.Query(lb, rb)
		if err != nil {
			t.Log(err)
			continue
		}
		sum2 := 0
		for j := lb; j <= rb; j++  {
			sum2 = sum2 + arr[j]
		}
		if sum1 != sum2 {
			t.Error("查询结果错误")
		}
	}
}
