package sort

import (
	"strconv"
)

//时间复杂度:平均:O(n*k)、最好:O(n*k)、最坏:O(n*k)
//空间复杂度:O(n+k)
//稳定性:稳定

/*
RadixSort
@Desc:基数排序
@Param:arr  []int32		待排数据
@Return: 	[]int32		排序结果
*/
func RadixSort(arr []int32) []int32 {
	max, _ := getMaxMinNum(arr)
	maxLen := len(strconv.Itoa(int(max)))
	divisor := int32(1) //被除数从1,10,100,1000.....

	//N次排序
	for i := 0; i < maxLen; i++ {
		//桶内值分配
		bucket := make([][]int32, 10)
		for _, value := range arr {
			index := value / divisor % 10
			bucket[index] = append(bucket[index], value)
		}

		//出桶重新组合
		tempArr := make([]int32, 0)
		for _, value := range bucket {
			for j := 0; j < len(value); j++ {
				tempArr = append(tempArr, value[j])
			}
		}

		arr = tempArr
		divisor *= 10
	}

	return arr
}
