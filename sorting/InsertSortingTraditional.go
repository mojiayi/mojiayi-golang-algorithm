package sorting

type InsertSortingTraditional struct{}

/**
* 传统的插入排序
 */
func (s *InsertSortingTraditional) Sort(arr *[]int) []int {
	var length = len(*arr)
	var sortedArr = make([]int, length)
	for index, value := range *arr {
		if index == 0 {
			// 第一个元素，不需要比较，直接放在已排序列表末尾
			sortedArr[length-1] = value
			continue
		}
		var insertPoint = -1
		var sortedStartPoint = length - index - 1
		// 从已排序列表末尾开始，从后往前比较，找到第一个比待排序数值小的，就找到了要被插入的下标
		for n := length - 1; n >= sortedStartPoint; n-- {
			if sortedArr[n] < value {
				insertPoint = n
				break
			}
		}

		var copySize = insertPoint - sortedStartPoint
		if copySize > 0 {
			// 要被移动的元素个数大于0时，才需要复制一遍数据
			for m := length - index; m <= insertPoint; m++ {
				sortedArr[m-1] = sortedArr[m]
			}
		}
		if insertPoint >= 0 {
			sortedArr[insertPoint] = value
		}
	}

	return sortedArr
}
