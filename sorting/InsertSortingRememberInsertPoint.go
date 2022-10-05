package sorting

type InsertSortingRememberInsertPoint struct{}

/**
* 记录上次插入点下标的插入排序，下一轮先判断新的待排序数值大于还是小于上次排序的数值
 */
func (s *InsertSortingRememberInsertPoint) Sort(arr *[]int) []int {
	var length = len(*arr)
	var sortedArr = make([]int, length)
	var insertPoint = -1
	for index, value := range *arr {
		if index == 0 {
			// 第一个元素，不需要比较，直接放在已排序列表末尾
			sortedArr[length-1] = value
			insertPoint = length - 1
			continue
		}

		var sortedStartPoint = length - index - 1
		var newInsertPoint = -1
		if value >= sortedArr[insertPoint] {
			// 如果待排序元素的数值大于或等于上一个被排序的，从上次的插入点往后查找新插入点
			for n := insertPoint; n < length; n++ {
				// 找到第一个大于待排序元素后，新插入点是这个元素的前一个位置
				if sortedArr[n] > value {
					newInsertPoint = n - 1
					break
				}
			}
			// 如果找到最后还没有大于待排序元素的，列表末尾就是新插入点
			if newInsertPoint < 0 {
				newInsertPoint = length - 1
			}
			insertPoint = newInsertPoint
		} else {
			// 如果待排序元素的数值小于上一个被排序的，从上次的插入点往前查找新插入点
			for n := insertPoint; n > sortedStartPoint-1; n-- {
				// 找到第一个小于待排序元素后，新插入点就是这个元素所在位置
				if sortedArr[n] < value {
					newInsertPoint = n
					break
				}
			}
			// 如果找到最前面还没有小于待排序元素的，则插入点是已排序下标的前一个位置
			if newInsertPoint < 0 {
				newInsertPoint = sortedStartPoint - 1
			}
			insertPoint = newInsertPoint
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
