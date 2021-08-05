package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	arr := []int32{930, 83, 63, 184, 505, 278, 8, 589, 109, -269}
	arr = sort.RadixSort(arr)
	fmt.Println(arr)
}
