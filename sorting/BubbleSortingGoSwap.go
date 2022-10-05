package sorting

type BubbleSortingGoSwap struct{}

/**
* 使用golang特有数据交换语法实现数值交换的冒泡排序
 */
func (s *BubbleSortingGoSwap) Sort(arr *[]int) []int {
	size := len(*arr)
	var i, j int
	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
	return *arr
}
