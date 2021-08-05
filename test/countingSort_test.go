package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	//最为基础的
	arr := []int32{2, 1, 10, 3, 7, 5, 10}
	res := sort.CountingSort(arr)
	fmt.Println(res)

	//优化后的
	arr = []int32{91, 99, 93, 95, 96, 94, 97, 91, 90}
	res = sort.CountingSortOptimize(arr)
	fmt.Println(res)

	//排序稳定版本
	arrModel := GetDataTest()
	resModel := sort.CountingSortStable(arrModel)
	fmt.Println(resModel)
}

func GetDataTest() []sort.CountingSortModel {
	return []sort.CountingSortModel{
		{
			Name: "张三",
			Sore: 91,
		},
		{
			Name: "李四",
			Sore: 99,
		},
		{
			Name: "王五",
			Sore: 95,
		},
		{
			Name: "赵六",
			Sore: 94,
		},
		{
			Name: "孙七",
			Sore: 95,
		},
	}
}
