package timewheel

import (
	"errors"
	"fmt"
	"mojiayi-golang-algorithm/domain"
	"mojiayi-golang-algorithm/linkedlist"
	"strconv"
	"time"
)

type RoundTimeWheel struct {
	domain.MojiayiTimeWheel
}

type RoundTaskNode struct {
	domain.TaskNode
	/**
	* 属于同一节点的任务
	 */
	TaskDetailList *[]RoundTaskDetail
}

type RoundTaskDetail struct {
	domain.TaskDetail
	/**
	* 任务轮次
	 */
	Round int
	/**
	* 创建时间，单位为秒
	 */
	PreviousRoundTime int
}

/**
* 创建一个刻度范围只有60秒的分轮次时间轮
 */
func (s *RoundTimeWheel) New() (*RoundTimeWheel, error) {
	now := time.Now().UnixMilli() / int64(domain.ONE_THOUSAND)
	roundTimeWheel := new(RoundTimeWheel)
	roundTimeWheel.MaxScale = domain.MAX_SCALE
	roundTimeWheel.CurrentScale = 0
	roundTimeWheel.StartupTime = int(now)
	roundTimeWheel.CurrentTime = int(now)
	roundTimeWheel.TaskNodeList = &linkedlist.CircleLinkedList{}
	for scale := 0; scale < domain.MAX_SCALE; scale++ {
		taskDetailList := make([]RoundTaskDetail, 0, 16)
		taskNode := RoundTaskNode{}
		taskNode.ID = scale + 1
		taskNode.Scale = scale
		taskNode.TaskDetailList = &taskDetailList
		roundTimeWheel.TaskNodeList.AddToTail(&linkedlist.Node{ID: scale, Data: taskNode})
	}
	return roundTimeWheel, nil
}

/**
* 添加只执行一次，不重复执行的任务
 */
func (s *RoundTimeWheel) AddOnceTask(delay int) error {
	return s.AddTask(delay, false)
}

/**
* 添加新任务，可指定是否重复执行
 */
func (s *RoundTimeWheel) AddTask(delay int, repeatFlag bool) error {
	round, scale, err := s.calculateRoundAndScale(delay)
	if err != nil {
		return err
	}

	taskNode, err := s.findNodeByRoundAndScale(round, scale)
	if err != nil {
		return err
	}

	id := strconv.Itoa(taskNode.ID) + "-" + strconv.Itoa(len(*taskNode.TaskDetailList)+1)
	newTask := RoundTaskDetail{}
	newTask.ID = id
	newTask.Scale = scale
	newTask.Round = round
	newTask.Delay = delay
	newTask.RepeatFlag = repeatFlag

	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	newTask.CreateTime = now
	newTask.PreviousRoundTime = now

	*taskNode.TaskDetailList = append(*taskNode.TaskDetailList, newTask)

	return nil
}

/**
* 删除任务
 */
func (s *RoundTimeWheel) DeleteTask(node *linkedlist.Node) (bool, error) {
	return true, nil
}

/**
* 调度并执行任务
 */
func (s *RoundTimeWheel) ExecuteTask() {
	node := s.TaskNodeList.Head

	for {
		now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
		taskList := *node.Data.(RoundTaskNode).TaskDetailList
		if len(taskList) > 0 {
			for index, task := range taskList {
				if task.Delay == 62 {
					fmt.Println("遇到62")
				}
				if task.Round == 0 {
					fmt.Println("执行任务(id=" + task.ID + ",scale=" + strconv.Itoa(task.Scale) + ",delay=" + strconv.Itoa(task.Delay) + ")")
				} else if task.Round > 0 {
					taskList[index].Round = task.Round - 1
					taskList[index].PreviousRoundTime = now
					roundTaskNode := RoundTaskNode{}
					roundTaskNode.ID = node.Data.(RoundTaskNode).ID
					roundTaskNode.Scale = node.Data.(RoundTaskNode).Scale
					roundTaskNode.TaskDetailList = &taskList
					(*node).Data = roundTaskNode
				}
			}
		}
		node = node.Next
		time.Sleep(time.Duration(1) * time.Second)
	}
}

/**
* 找到对应轮次和刻度的节点
 */
func (s *RoundTimeWheel) findNodeByRoundAndScale(round int, scale int) (RoundTaskNode, error) {
	node := s.TaskNodeList.Head.Next
	for node.ID != s.TaskNodeList.Head.ID {
		taskNode := node.Data.(RoundTaskNode)
		if taskNode.Scale == scale {
			return taskNode, nil
		}
		node = node.Next
	}
	return RoundTaskNode{}, errors.New("没有scale=" + strconv.Itoa(scale) + "的节点")
}

/**
* 计算新任务所属的轮次和刻度
 */
func (s *RoundTimeWheel) calculateRoundAndScale(delay int) (int, int, error) {
	round, scale := -1, -1
	if delay <= 0 {
		return round, scale, errors.New("延迟时间必须大于0")
	}
	round = delay / s.MaxScale
	scale = (delay + s.CurrentScale) % s.MaxScale
	return round, scale, nil
}
