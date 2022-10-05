package sorting

type SelectionSortingOnlyOneArray struct{}

/**
 * 不新建一个数组的选择排序算法实现方案
 *
 */
func (s *SelectionSortingOnlyOneArray) Sort(arr *[]int) []int {
	// 初始化一个理论上的最大值，因为本示例是以int为元素类型
	var MAX_VALUE = 2 ^ 32 - 1
	var length = len(*arr)
	for m := 0; m < length; m++ {
		// 初始化每轮外层循环的最小元素值和下标
		var minValueIndex = -1
		var minValue = MAX_VALUE
		for n := m; n < length; n++ {
			// 从未排序元素中查找最小元素值和下标
			if (*arr)[n] < minValue {
				minValue = (*arr)[n]
				minValueIndex = n
			}
		}
		if minValueIndex >= 0 {
			// 把未排序元素往后移动一个位置
			for k := minValueIndex; k > m; k-- {
				(*arr)[k] = (*arr)[k-1]
			}
			// 把本轮查找到的最小元素放到已排序元素末尾
			(*arr)[m] = minValue
		}
	}
	return *arr
}
