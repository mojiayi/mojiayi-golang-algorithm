package algorithm_test

import (
	"fmt"
	"mojiayi-golang-algorithm/linkedlist"
	"testing"
)

func TestCircleLinkedList(t *testing.T) {
	taskList := make([]string, 3)
	taskList[0] = "任务0"
	taskList[1] = "任务1"
	taskList[2] = "任务2"
	linkedList := linkedlist.CircleLinkedList{}
	node2 := &linkedlist.Node{ID: 2}
	node3 := &linkedlist.Node{ID: 3}
	node4 := &linkedlist.Node{ID: 4}
	node5 := &linkedlist.Node{ID: 5}
	node6 := &linkedlist.Node{ID: 6}
	node7 := &linkedlist.Node{ID: 7}
	node8 := &linkedlist.Node{ID: 8}
	linkedList.AddToHead(&linkedlist.Node{ID: 1, Data: taskList})
	linkedList.AddToTail(node2)
	linkedList.AddToTail(node4)
	linkedList.Add(node3, node2)
	linkedList.AddToTail(node5)
	linkedList.Add(node6, node5)
	linkedList.AddToTail(node8)
	linkedList.Add(node7, node6)

	fmt.Print("初始链表：")
	linkedList.Print()

	fmt.Print("删除链表头节点：")
	linkedList.DeleteHead()
	linkedList.Print()

	fmt.Print("删除链表尾节点：")
	linkedList.DeleteTail()
	linkedList.Print()

	fmt.Print("删除链表id=6的节点：")
	linkedList.Delete(node6)
	linkedList.Print()
}
