package sorting

type SelectionSortingTraditional struct{}

/**
 * 传统的选择排序算法实现方案，有两个for循环，使用2个数组
 *
 */
func (s *SelectionSortingTraditional) Sort(arr *[]int) []int {
	// 初始化一个理论上的最大值，因为本示例是以int为元素类型
	var MAX_VALUE = 2 ^ 32 - 1
	var MIN_VALUE = -2 ^ 32
	var length = len(*arr)
	// 创建一个容量与待排序数组相同的新数组
	var sortedArr = make([]int, length)
	for n := 0; n < length; n++ {
		// 初始化每轮外层循环的最小元素值和下标
		var minValueIndex = -1
		var minValue = MAX_VALUE
		for m := 0; m < length; m++ {
			// 从未排序元素中查找最小元素值和下标
			if (*arr)[m] > MIN_VALUE && (*arr)[m] < minValue {
				minValue = (*arr)[m]
				minValueIndex = m
			}
		}
		if minValueIndex >= 0 {
			// 把本轮查找到的最小元素放到已排序元素末尾
			sortedArr[n] = minValue
			(*arr)[minValueIndex] = MIN_VALUE
		}
	}
	return sortedArr
}
