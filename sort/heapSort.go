package sort

//时间复杂度:平均:O(nlogn)、最好:O(nlogn)、最坏:O(nlogn)
//空间复杂度:O(1)
//稳定性:稳定

/*
HeapSort
@Desc:堆排
@Param:arr 	[]int32		待排数据
@Return: 	[]int32		排序结果
*/
func HeapSort(arr []int32) []int32 {
	//先构建堆
	buildHeap(arr)

	//堆顶与末尾交换,堆顶再次堆化
	lastNode := len(arr) - 1
	for i := lastNode; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, int32(i), 0)
	}

	return arr
}

/*
heapify
@Desc:堆化
@Param:tree []int32 树,数组
@Param:n 	int32	树的元素个数
@Param:i 	int32	针对哪个节点进行堆化
*/
func heapify(tree []int32, n int32, i int32) {
	if i >= n {
		return
	}

	max := i         //定义i为parent节点
	left := 2*i + 1  //parent的左节点
	right := 2*i + 2 //parent的右节点

	//找最大元素
	if left < n && tree[left] > tree[max] {
		max = left
	}
	if right < n && tree[right] > tree[max] {
		max = right
	}

	//交换堆顶
	if max != i {
		tree[i], tree[max] = tree[max], tree[i]
		heapify(tree, n, max)
	}
}

/*
buildHeap
@Desc:构建堆
@Param:tree []int32 树,数组
*/
func buildHeap(tree []int32) {
	l := int32(len(tree))
	lastNode := l - 1                //最后一个节点
	parentOfLastNode := lastNode / 2 //最后一个节点的父节点

	//由下到上开始堆化构建堆
	for i := parentOfLastNode; i >= 0; i-- {
		heapify(tree, l, i)
	}
}
