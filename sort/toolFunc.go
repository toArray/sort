package sort

// CountingSortModel 稳定计数排序测试用例数值
type CountingSortModel struct {
	Name string
	Sore int32
}

/*
getMaxMinFloat64
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxMinFloat64(arr []float64) (max float64, min float64) {
	if len(arr) == 0 {
		return
	}

	max = arr[0]
	min = arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return
}

/*
getMaxMinOfModel
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxMinOfModel(arr []CountingSortModel) (max int32, min int32) {
	if len(arr) == 0 {
		return
	}

	max = arr[0].Sore
	min = arr[0].Sore

	for _, value := range arr {
		if value.Sore > max {
			max = value.Sore
		}
		if value.Sore < min {
			min = value.Sore
		}
	}

	return
}

/*
getMaxMinNum
@Desc:获得数组内最大值和最小值
@Param:arr 	[]int32	数组
@Return:max int32	最大值
@Return:min int32	最小值
*/
func getMaxMinNum(arr []int32) (max int32, min int32) {
	if len(arr) == 0 {
		return
	}

	max = arr[0]
	min = arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	return
}

/*
getMaxLenOfString
@Desc:获得最长的字符串长度
@Param:arr 	[]string	字符串数组
@Return: 	int32		最大长度
*/
func getMaxLenOfString(arr []string) int {
	maxLen := 0
	for _, value := range arr {
		l := len(value)
		if l >= maxLen {
			maxLen = l
		}
	}

	return maxLen
}
