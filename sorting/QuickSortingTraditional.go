package sorting

type QuickSortingTraditional struct{}

func (s QuickSortingTraditional) Sort(arr *[]int) []int {
	var length = len(*arr)
	if length == 0 {
		return *arr
	}

	// 1. 从数组头往后找大于基数的元素，从数组尾往前找小于基数的元素
	quickSort(arr, 0, length-1)

	return *arr
}

func quickSort(arr *[]int, leftIndex int, rightIndex int) {
	if leftIndex >= rightIndex {
		// 2. 如果左右两个下标值相遇，说明已经完成本轮递归
		return
	}
	var m = leftIndex
	var n = rightIndex
	// 3. 取数组的第一个元素作为基准值
	var benchmark = (*arr)[m]
	for {
		if m < n {
			for m < n && (*arr)[n] >= benchmark {
				// 4. 从右往左查找大于等于基准值的元素
				n--
			}
			if m < n {
				// 5. 找到大于等于基准值的元素后，把这个元素填充到取得基准值的位置
				(*arr)[m] = (*arr)[n]
				// 6. 左起查找的下标值+1，不再比较刚刚被重新设置的值
				m++
			}
			for m < n && (*arr)[m] < benchmark {
				// 7. 从左往右查找小于基准值的元素
				m++
			}
			if m < n {
				// 8. 找到小于基准值的元素后，把这个元素值填充到第5步里被迁移了的元素位置
				(*arr)[n] = (*arr)[m]
				// 9. 右起查找的下标值-1，不再比较刚刚被重新设置的值
				n--
			}
		} else {
			break
		}
	}
	// 10. 把基准值填入第8步里被迁移了的元素位置
	(*arr)[m] = benchmark
	// 11. 对本轮基准元素左侧的元素排序
	quickSort(arr, leftIndex, m-1)
	// 12. 对本轮基准元素右侧的元素排序
	quickSort(arr, m+1, rightIndex)
}
