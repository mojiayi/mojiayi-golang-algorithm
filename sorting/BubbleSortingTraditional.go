package sorting

type BubbleSortingTraditional struct{}

/**
* 传统的冒泡排序，用一个临时变量实现数值交换
 */
func (s BubbleSortingTraditional) Sort(arr *[]int) []int {
	size := len(*arr)
	var i, j int
	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			if (*arr)[i] > (*arr)[j] {
				tmp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = tmp
			}
		}
	}
	return *arr
}
