package sorting

type SelectionSortingOnlyOneArray struct{}

/**
 * 不新建一个数组的选择排序算法实现方案
 *
 */
func (s SelectionSortingOnlyOneArray) Sort(arr *[]int) []int {
	var MAX_VALUE = 2147483647
	var length = len(*arr)
	for m := 0; m < length; m++ {
		var minValueIndex = -1
		var minValue = MAX_VALUE
		for n := m; n < length; n++ {
			if (*arr)[n] < minValue {
				minValue = (*arr)[n]
				minValueIndex = n
			}
		}
		if minValueIndex >= 0 {
			for k := minValueIndex; k > m; k-- {
				(*arr)[k] = (*arr)[k-1]
			}
			(*arr)[m] = minValue
		}
	}
	return *arr
}
