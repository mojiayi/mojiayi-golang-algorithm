package sorting_test

import (
	sorting "mojiayi-golang-algorithm/sorting"
	"testing"
)

var arr = []int{4, 7, 2, 3, 9}

func TestBubbleSortingExclusiveOr(t *testing.T) {
	var bubbleSorting sorting.BubbleSortingExclusiveOr
	executeBubbleSort(bubbleSorting, t)
}

func TestBubbleSortingGoSwap(t *testing.T) {
	var bubbleSorting sorting.BubbleSortingGoSwap
	executeBubbleSort(bubbleSorting, t)
}

func TestBubbleSortingPlusMinus(t *testing.T) {
	var bubbleSorting sorting.BubbleSortingPlusMinus
	executeBubbleSort(bubbleSorting, t)
}

func TestBubbleSortingTraditional(t *testing.T) {
	var bubbleSorting sorting.BubbleSortingTraditional
	executeBubbleSort(bubbleSorting, t)
}

func TestInsertSortingTraditional(t *testing.T) {
	var insertSorting sorting.InsertSortingTraditional
	executeInsertSort(insertSorting, t)
}

func TestInsertSortingRememberInsertPoint(t *testing.T) {
	var insertSorting sorting.InsertSortingRememberInsertPoint
	executeInsertSort(insertSorting, t)
}

func TestSelectionSortingTraditional(t *testing.T) {
	var selectionSorting sorting.SelectionSortingTraditional
	executeInsertSort(selectionSorting, t)
}

func TestSelectionSortingOnlyOneArray(t *testing.T) {
	var selectionSorting sorting.SelectionSortingOnlyOneArray
	executeInsertSort(selectionSorting, t)
}

func TestQuickSortingTraditional(t *testing.T) {
	var quickSorting sorting.QuickSortingTraditional

	executeInsertSort(quickSorting, t)
}

/**
* 执行指定的排序操作，并校验排序结果是否正确
 */
func executeBubbleSort(sortExecutor sorting.ISorting, t *testing.T) {
	sortExecutor.Sort(&arr)

	checkResult(0, 2, arr, t)
	checkResult(3, 7, arr, t)
	checkResult(4, 9, arr, t)
}

func executeInsertSort(sortExecutor sorting.ISorting, t *testing.T) {
	sortedArr := sortExecutor.Sort(&arr)

	checkResult(0, 2, sortedArr, t)
	checkResult(3, 7, sortedArr, t)
	checkResult(4, 9, sortedArr, t)
}

func checkResult(index int, expected int, arr []int, t *testing.T) {
	if arr[index] != expected {
		t.Errorf("sortedArr[%v] expected:%v,actual:%v", index, expected, arr[index])
	}
}
