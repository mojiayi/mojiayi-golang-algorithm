package algorithm_test

import (
	"mojiayi-golang-algorithm/timewheel"
	"testing"
)

func TestSimpleTaskNodeLinkedList(t *testing.T) {
	taskList := make([]string, 3)
	taskList[0] = "任务0"
	taskList[1] = "任务1"
	taskList[2] = "任务2"
	linkedList := timewheel.CircleLinkedList{}
	linkedList.AddToHead(&timewheel.Node{ID: 1, Data: taskList})
	linkedList.AddToTail(&timewheel.Node{ID: 2})
	linkedList.AddToTail(&timewheel.Node{ID: 3})
	linkedList.AddToTail(&timewheel.Node{ID: 4})
	linkedList.AddToTail(&timewheel.Node{ID: 5})
	linkedList.Print()

	linkedList.DeleteHead()
	linkedList.Print()

	linkedList.DeleteTail()
	linkedList.Print()
}
