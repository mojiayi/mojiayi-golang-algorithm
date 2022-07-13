package sorting

type BubbleSortingPlusMinus struct{}

/**
* 使用两个数值相加相减实现数值交换的冒泡排序
 */
func (s BubbleSortingPlusMinus) Sort(arr *[]int) []int {
	size := len(*arr)
	var i, j int
	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i] += (*arr)[j]
				(*arr)[j] = (*arr)[i] - (*arr)[j]
				(*arr)[i] = (*arr)[i] - (*arr)[j]
			}
		}
	}
	return *arr
}
