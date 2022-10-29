package timewheel

import (
	"errors"
	"fmt"
	"mojiayi-golang-algorithm/domain"
	"mojiayi-golang-algorithm/linkedlist"
	"strconv"
	"sync"
	"time"
)

/**
* 添加了轮次概念的时间轮
 */
type RoundTimeWheel struct {
	domain.MojiayiTimeWheel
}

/**
* 轮次时间轮的节点
 */
type RoundTaskNode struct {
	domain.TaskNode
	/**
	* 属于同一节点的任务
	 */
	TaskDetailList *[]RoundTaskDetail
}

/**
* 轮次时间轮的任务详情
 */
type RoundTaskDetail struct {
	domain.TaskDetail
	/**
	* 任务轮次，任务创建后固定不变
	 */
	Round int
	/**
	* 当前轮次，调度过程中递增
	 */
	currentRound int
	/**
	* 当前刻度，单位为秒，调度过程中递增
	 */
	currentScale int
}

/**
* 创建一个刻度范围只有60秒的轮次时间轮
 */
func (s *RoundTimeWheel) New() (*RoundTimeWheel, error) {
	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	roundTimeWheel := new(RoundTimeWheel)
	roundTimeWheel.MaxScale = domain.MAX_SCALE
	roundTimeWheel.CurrentScale = 0
	roundTimeWheel.StartupTime = now
	roundTimeWheel.CurrentTime = now
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

	taskNode, err := s.findNodeByScale(scale)
	if err != nil {
		return err
	}

	id := strconv.Itoa(taskNode.ID) + "-" + strconv.Itoa(len(*taskNode.TaskDetailList)+1)
	newTask := RoundTaskDetail{}
	newTask.ID = id
	newTask.Scale = scale
	newTask.Round = round
	newTask.currentRound = 0
	newTask.currentScale = 0
	newTask.Delay = delay
	newTask.RepeatFlag = repeatFlag

	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	newTask.CreateTime = now

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
	var wg sync.WaitGroup
	node := s.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for node.Next.ID != s.TaskNodeList.Head.ID {
			taskList := node.Data.(RoundTaskNode).TaskDetailList
			if len(*taskList) > 0 {
				wg.Add(1)
				go func() {
					s.executeSameScaleTask(taskList)
					defer wg.Done()
				}()
			}
			node = node.Next
		}
		wg.Wait()
		node = s.TaskNodeList.Head
	}
}

func (s *RoundTimeWheel) executeSameScaleTask(taskList *[]RoundTaskDetail) {
	for index, task := range *taskList {
		if task.ExecuteFlag {
			continue
		}
		runnable := s.isRunnable(task)
		if runnable {
			fmt.Print("执行轮次时间轮任务(id=" + task.ID + ",scale=" + strconv.Itoa(task.Scale) + ",delay=" + strconv.Itoa(task.Delay) + ")，")
			fmt.Println("时间=" + time.Now().Format("2006-01-02 15:04:05"))
			(*taskList)[index].ExecuteFlag = true
			continue
		}

		currentRound := (*taskList)[index].currentRound
		currentScale := (*taskList)[index].currentScale

		addRoundFlag := false
		currentScale += 1
		if currentScale == s.MaxScale {
			currentScale = 0
			addRoundFlag = true
		}

		if addRoundFlag {
			currentRound += 1
			if currentRound >= task.Round {
				currentRound = task.Round
			}
		}

		// fmt.Println("轮次时间轮任务(id=" + task.ID + ")最新计时round=" + strconv.Itoa(currentRound) + ",scale=" + strconv.Itoa(currentScale))
		(*taskList)[index].currentRound = currentRound
		(*taskList)[index].currentScale = currentScale
	}
}

/**
* 判断任务是否达到运行条件
 */
func (s *RoundTimeWheel) isRunnable(taskDetail RoundTaskDetail) bool {
	if taskDetail.ExecuteFlag {
		return false
	}
	if taskDetail.currentRound != taskDetail.Round {
		return false
	}
	if taskDetail.currentScale != taskDetail.Scale {
		return false
	}
	return true
}

/**
* 找到对应刻度的节点
 */
func (s *RoundTimeWheel) findNodeByScale(scale int) (RoundTaskNode, error) {
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
