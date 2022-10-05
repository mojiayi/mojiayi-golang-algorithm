package sorting

type BubbleSortingExclusiveOr struct{}

/**
* 使用异或操作实现数值交换的冒泡排序
 */
func (s *BubbleSortingExclusiveOr) Sort(arr *[]int) []int {
	size := len(*arr)
	var i, j int
	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i] ^= (*arr)[j]
				(*arr)[j] ^= (*arr)[i]
				(*arr)[i] ^= (*arr)[j]
			}
		}
	}
	return *arr
}
