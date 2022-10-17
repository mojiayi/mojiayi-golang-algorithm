package algorithm_test

import (
	"mojiayi-golang-algorithm/linkedlist"
	"mojiayi-golang-algorithm/timewheel"
	"testing"
)

func TestCircleLinkedList(t *testing.T) {
	taskList := make([]string, 3)
	taskList[0] = "任务0"
	taskList[1] = "任务1"
	taskList[2] = "任务2"
	linkedList := linkedlist.CircleLinkedList{}
	linkedList.AddToHead(&linkedlist.Node{ID: 1, Data: taskList})
	linkedList.AddToTail(&linkedlist.Node{ID: 2})
	linkedList.AddToTail(&linkedlist.Node{ID: 3})
	linkedList.AddToTail(&linkedlist.Node{ID: 4})
	linkedList.AddToTail(&linkedlist.Node{ID: 5})
	linkedList.Print()

	linkedList.DeleteHead()
	linkedList.Print()

	linkedList.DeleteTail()
	linkedList.Print()
}

func TestSimpleTimeWheel(t *testing.T) {
	simpleTimeWheel := timewheel.SimpleTimeWheel{}
	instance, err := simpleTimeWheel.New()
	if err != nil {
		t.Errorf("构建简单时间轮失败%v", err)
	}
	instance.AddOnceTask(1)
	instance.AddOnceTask(1)
	instance.AddOnceTask(3)
	instance.AddOnceTask(3)
	instance.AddOnceTask(4)
	instance.AddOnceTask(5)
	instance.AddOnceTask(6)
	instance.AddOnceTask(7)
	instance.AddOnceTask(8)
	instance.AddOnceTask(9)
	instance.AddOnceTask(10)
	instance.AddOnceTask(10)
	instance.AddOnceTask(10)
	instance.AddOnceTask(20)
	instance.ExecuteTask()
}

func TestRoundTimeWheel(t *testing.T) {
	roundTimeWheel := timewheel.RoundTimeWheel{}
	instance, err := roundTimeWheel.New()
	if err != nil {
		t.Errorf("构建轮次时间轮失败%v", err)
	}
	instance.AddOnceTask(1)
	instance.AddOnceTask(1)
	instance.AddOnceTask(2)
	instance.AddOnceTask(3)
	instance.AddOnceTask(3)
	instance.AddOnceTask(62)
	instance.AddOnceTask(4)
	instance.AddOnceTask(5)
	instance.AddOnceTask(6)
	instance.AddOnceTask(7)
	instance.AddOnceTask(8)
	instance.AddOnceTask(9)
	instance.AddOnceTask(10)
	instance.AddOnceTask(10)
	instance.AddOnceTask(10)
	instance.AddOnceTask(20)
	instance.ExecuteTask()
}
