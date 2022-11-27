package timewheel

import (
	"errors"
	"fmt"
	"mojiayi-golang-algorithm/domain"
	"mojiayi-golang-algorithm/linkedlist"
	"strconv"
	"time"
)

/**
* 简单时间轮
 */
type SimpleTimeWheel struct {
	domain.MojiayiTimeWheel
}

/**
* 简单时间轮的节点
 */
type SimpleTaskNode struct {
	domain.TaskNode
	/**
	* 属于同一节点的任务
	 */
	TaskDetailList *[]domain.TaskDetail
}

/**
* 创建一个刻度范围只有60秒的简单时间轮
 */
func (s *SimpleTimeWheel) New() (*SimpleTimeWheel, error) {
	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	simpleTimeWheel := new(SimpleTimeWheel)
	simpleTimeWheel.MaxScale = domain.SECOND_SCALE
	simpleTimeWheel.CurrentScale = 0
	simpleTimeWheel.StartupTime = now
	simpleTimeWheel.CurrentTime = now
	simpleTimeWheel.TaskNodeList = &linkedlist.CircleLinkedList{}
	for scale := 0; scale < domain.SECOND_SCALE; scale++ {
		taskDetailList := make([]domain.TaskDetail, 0, 16)
		taskNode := SimpleTaskNode{}
		taskNode.ID = scale + 1
		taskNode.Scale = scale
		taskNode.TaskDetailList = &taskDetailList
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
	newTask := domain.TaskDetail{ID: id, Scale: scale, Delay: delay, RepeatFlag: repeatFlag}

	*taskNode.TaskDetailList = append(*taskNode.TaskDetailList, newTask)

	return nil
}

/**
* 调度并执行任务
 */
func (s *SimpleTimeWheel) ExecuteTask() {
	node := s.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		taskList := node.Data.(SimpleTaskNode).TaskDetailList
		if len(*taskList) > 0 {
			for _, task := range *taskList {
				if task.ExecuteFlag {
					continue
				}
				fmt.Print("执行简单时间轮任务(id=" + task.ID + ",scale=" + strconv.Itoa(task.Scale) + ")，")
				fmt.Println("时间=" + time.Now().Format("2006-01-02 15:04:05"))
			}
		}
		node = node.Next
	}
	ticker.Stop()
}

/**
* 找到对应刻度的节点
 */
func (s *SimpleTimeWheel) findNodeByScale(scale int) (*SimpleTaskNode, error) {
	node := s.TaskNodeList.Head.Next
	for node.ID != s.TaskNodeList.Head.ID {
		taskNode := node.Data.(SimpleTaskNode)
		if taskNode.Scale == scale {
			return &taskNode, nil
		}
		node = node.Next
	}
	return nil, errors.New("没有scale=" + strconv.Itoa(scale) + "的节点")
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
