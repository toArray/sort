package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

//希尔排序
func TestShellSort(t *testing.T) {

	sortData := []int32{9, 2, 11, 4, 65, 8, 1, 1, 9, 9, 3}

	res := sort.ShellSort(sortData)

	fmt.Println(res)
}
