package test

import (
	"fmt"
	"luoqiangSort/sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	arr := []float64{2.4, 2.32, 3.1, 4.37, 5.7, 6.01, 10}
	arr = sort.BucketSort(arr, 0)
	fmt.Println(arr)
}
