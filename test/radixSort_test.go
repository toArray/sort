package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	arr := []int32{930, 83, 63, 184, 505, 278, 8, 589, 109, 269, 12, 4, 7, 123, 457, 124, 8, 24}
	arr = sort.RadixSort(arr)
	fmt.Println(arr)
}

func TestRadixSortWord(t *testing.T) {
	arr := []string{"ab1", "dddd", "asdas", "fdhdh", "abc", "abcd", "zzasdas", "kdasd", "1226", "234"}
	arr = sort.RadixSortOfWord(arr)
	fmt.Println(arr)
}
