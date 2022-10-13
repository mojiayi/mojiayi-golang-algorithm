package algorithm_test

import (
	"mojiayi-golang-algorithm/timewheel"
	"testing"
)

func TestSimpleTaskNodeLinkedList(t *testing.T) {
	linkedList := timewheel.CircleLinkedList{}

	linkedList.AddToHead(&timewheel.SimpleTaskNode{ID: 1, Next: nil})
	linkedList.AddToTail(&timewheel.SimpleTaskNode{ID: 2})
	linkedList.AddToTail(&timewheel.SimpleTaskNode{ID: 3})
	linkedList.AddToTail(&timewheel.SimpleTaskNode{ID: 4})
	linkedList.AddToTail(&timewheel.SimpleTaskNode{ID: 5})
	linkedList.Print()

	linkedList.DeleteHead()
	linkedList.Print()

	linkedList.DeleteTail()
	linkedList.Print()
}
