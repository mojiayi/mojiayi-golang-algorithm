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
	fmt.Print("添加第一节点：")
	linkedList.Print()

	linkedList.AddToTail(node2)
	fmt.Print("添加新的尾节点2：")
	linkedList.Print()

	linkedList.AddToTail(node4)
	fmt.Print("添加新的尾节点4：")
	linkedList.Print()

	linkedList.Add(node3, node2)
	fmt.Print("添加新节点3到2之后：")
	linkedList.Print()

	linkedList.AddToTail(node5)
	fmt.Print("添加新的尾节点5：")
	linkedList.Print()

	linkedList.Add(node6, node5)
	fmt.Print("添加新节点6到5之后：")
	linkedList.Print()

	linkedList.AddToTail(node8)
	fmt.Print("添加新的尾节点8：")
	linkedList.Print()

	linkedList.Add(node7, node6)
	fmt.Print("添加新节点7到6之后：")
	linkedList.Print()

	fmt.Print("完成构建本测试用例的链表：")
	linkedList.Print()

	fmt.Print("删除链表头节点1：")
	linkedList.DeleteHead()
	linkedList.Print()

	fmt.Print("删除链表尾节点8：")
	linkedList.DeleteTail()
	linkedList.Print()

	fmt.Print("删除链表id=6的节点：")
	linkedList.Delete(node6)
	linkedList.Print()
}
