package timewheel

import (
	"errors"
	"fmt"
	"mojiayi-golang-algorithm/linkedlist"
	"strconv"
	"time"
)

type SimpleTimeWheel struct {
	/**
	* 最大刻度值，单位为秒
	 */
	MaxScale int
	/**
	* 当前刻度值，单位为秒
	 */
	CurrentScale int
	/**
	* 启动时间，单位为秒
	 */
	StartupTime int
	/**
	* 当前时间，单位为秒
	 */
	CurrentTime int
	/**
	* 任务列表
	 */
	TaskNodeList *linkedlist.CircleLinkedList
}

type TaskNode struct {
	/**
	* 任务节点id
	 */
	ID int
	/**
	* 所属刻度，单位为秒
	 */
	Scale int
	/**
	* 属于同一节点的任务
	 */
	TaskDetailList *[]TaskDetail
}

type TaskDetail struct {
	/**
	* 任务id
	 */
	ID string
	/**
	* 延迟时间，单位为秒
	 */
	Delay int
	/**
	* 所属刻度，单位为秒
	 */
	Scale int
	/**
	* 是否重复执行
	 */
	RepeatFlag bool
}

var (
	MAX_SCALE int = 60
)

/**
* 创建一个刻度范围只有60秒的简单时间轮
 */
func (s *SimpleTimeWheel) New() (*SimpleTimeWheel, error) {
	now := time.Second
	simpleTimeWheel := new(SimpleTimeWheel)
	simpleTimeWheel.MaxScale = MAX_SCALE
	simpleTimeWheel.CurrentScale = 0
	simpleTimeWheel.StartupTime = int(now)
	simpleTimeWheel.CurrentTime = int(now)
	simpleTimeWheel.TaskNodeList = &linkedlist.CircleLinkedList{}
	for scale := 0; scale < MAX_SCALE; scale++ {
		taskDetailList := make([]TaskDetail, 0, 16)
		taskNode := TaskNode{ID: scale, Scale: scale, TaskDetailList: &taskDetailList}
		simpleTimeWheel.TaskNodeList.AddToTail(&linkedlist.Node{ID: scale, Data: taskNode})
	}
	return simpleTimeWheel, nil
}

/**
* 添加只执行一次，不重复执行的任务
 */
func (s *SimpleTimeWheel) AddOnceTask(delay int) error {
	return s.AddTask(delay, false)
}

/**
* 添加新任务，可指定是否重复执行
 */
func (s *SimpleTimeWheel) AddTask(delay int, repeatFlag bool) error {
	scale, err := s.calculateScale(delay)
	if err != nil {
		return err
	}

	taskNode, err := s.findNodeByScale(scale)
	if err != nil {
		return err
	}

	id := strconv.Itoa(taskNode.ID) + "-" + strconv.Itoa(len(*taskNode.TaskDetailList)+1)
	newTask := TaskDetail{ID: id, Scale: scale, Delay: delay, RepeatFlag: repeatFlag}

	*taskNode.TaskDetailList = append(*taskNode.TaskDetailList, newTask)

	return nil
}

/**
* 删除任务
 */
func (s *SimpleTimeWheel) DeleteTask(node *linkedlist.Node) (bool, error) {
	return true, nil
}

/**
* 调度并执行任务
 */
func (s *SimpleTimeWheel) ExecuteTask() {
	node := s.TaskNodeList.Head

	for {
		taskList := node.Data.(TaskNode).TaskDetailList
		if len(*taskList) > 0 {
			for _, task := range *taskList {
				fmt.Println("执行任务(id=" + task.ID + ",scale=" + strconv.Itoa(task.Scale) + ")")
			}
		}
		node = node.Next
		time.Sleep(time.Duration(1) * time.Second)
	}
}

/**
* 找到对应刻度的节点
 */
func (s *SimpleTimeWheel) findNodeByScale(scale int) (TaskNode, error) {
	node := s.TaskNodeList.Head.Next
	for node.ID != s.TaskNodeList.Head.ID {
		taskNode := node.Data.(TaskNode)
		if taskNode.Scale == scale {
			return taskNode, nil
		}
		node = node.Next
	}
	return TaskNode{}, errors.New("没有scale=" + strconv.Itoa(scale) + "的节点")
}

/**
* 计算新任务所属的刻度
 */
func (s *SimpleTimeWheel) calculateScale(delay int) (int, error) {
	if delay <= 0 {
		return -1, errors.New("延迟时间必须大于0")
	}
	if delay > s.MaxScale {
		return -1, errors.New("延迟时间不能超过" + strconv.Itoa(s.MaxScale))
	}
	scale := (delay + s.CurrentScale) % s.MaxScale
	return scale, nil
}
