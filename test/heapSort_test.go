package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int32{2, 1, 10, 3, 7, 5, 10}
	sort.HeapSort(arr)
	fmt.Println(arr)
}

//func TestBuildHeap(t *testing.T) {
//	arr := []int32{2, 1, 10, 3, 7, 5, 10}
//	sort.buildHeap(arr)
//	fmt.Println(arr)
//}
//
//func TestHeapify(t *testing.T) {
//	arr := []int32{2, 1, 10, 3, 7, 5, 10}
//	sort.heapify(arr, int32(len(arr)), 0)
//	fmt.Println(arr)
//}
