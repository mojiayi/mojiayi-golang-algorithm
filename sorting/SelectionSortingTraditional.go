package sorting

type SelectionSortingTraditional struct{}

/**
 * 传统的选择排序算法实现方案，有两个for循环，使用2个数组
 *
 */
func (s SelectionSortingTraditional) Sort(arr *[]int) []int {
	var MAX_VALUE = 2147483647
	var MIN_VALUE = -2147483648
	var length = len(*arr)
	var sortedArr = make([]int, length)
	for n := 0; n < length; n++ {
		var minValueIndex = -1
		var minValue = MAX_VALUE
		for m := 0; m < length; m++ {
			if (*arr)[m] > MIN_VALUE && (*arr)[m] < minValue {
				minValue = (*arr)[m]
				minValueIndex = m
			}
		}
		if minValueIndex >= 0 {
			sortedArr[n] = minValue
			(*arr)[minValueIndex] = MIN_VALUE
		}
	}
	return sortedArr
}
