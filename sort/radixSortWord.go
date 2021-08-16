package sort

//时间复杂度:平均:O(n*k)、最好:O(n*k)、最坏:O(n*k)
//空间复杂度:O(n+k)
//稳定性:稳定

/*
RadixSortOfWord
@Desc:基数排序 - 单词
@Param:arr  []int32		待排数据
@Return: 	[]int32		排序结果
*/
func RadixSortOfWord(arr []string) []string {
	maxLen := getMaxLenOfString(arr)
	divisor := int32(maxLen - 1)

	//N次排序
	for i := 0; i < maxLen; i++ {
		//桶内值分配
		bucket := make([][]string, 128)
		for _, value := range arr {
			b := []byte(value)
			index := int32(0)
			if int32(len(b)) > divisor {
				index = int32(b[divisor])
			}

			bucket[index] = append(bucket[index], value)
		}

		//出桶重新组合
		tempArr := make([]string, 0)
		for _, value := range bucket {
			for j := 0; j < len(value); j++ {
				tempArr = append(tempArr, value[j])
			}
		}

		arr = tempArr
		divisor--
	}

	return arr
}
