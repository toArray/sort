package test

import (
	"fmt"
	"luoqiangSort/sort"
	"math/rand"
	"testing"
	"time"
)

//插入排序
func TestMergeSort(t *testing.T) {
	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		sortData = append(sortData, int32(rand.Intn(100)))
	}

	res := sort.MergeSort(sortData)

	fmt.Println(res)
}
