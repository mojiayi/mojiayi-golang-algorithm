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
* 分层时间轮
 */
type MultiLayerTimeWheel struct {
	domain.MojiayiTimeWheel
}

/**
* 分层时间轮的任务详情
 */
type MultiLayerTaskDetail struct {
	domain.TaskDetail
	/**
	* 时针层刻度，任务创建后固定不变
	 */
	Hour int
	/**
	* 分针层刻度，任务创建后固定不变
	 */
	Minute int
	/**
	* 秒针层刻度，任务创建后固定不变
	 */
	Second int
	/**
	* 时针层当前刻度，调度过程中递增
	 */
	currentHour int
	/**
	 * 分针层当前刻度，调度过程中递增
	 */
	currentMinute int
	/**
	 * 秒针层当前刻度，调度过程中递增
	 */
	currentSecond int
}

/**
* 分层时间轮节点
 */
type MultiLayerTaskNode struct {
	domain.TaskNode
	/**
	* 属于同一节点的任务
	 */
	TaskDetailList *[]MultiLayerTaskDetail
}

/**
* 创建一个刻度范围只有60秒的分轮次时间轮
 */
func (s *MultiLayerTimeWheel) New() (*MultiLayerTimeWheel, error) {
	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	multiLayerTimeWheel := new(MultiLayerTimeWheel)
	multiLayerTimeWheel.MaxScale = domain.SECOND_SCALE
	multiLayerTimeWheel.CurrentScale = 0
	multiLayerTimeWheel.StartupTime = now
	multiLayerTimeWheel.CurrentTime = now
	multiLayerTimeWheel.TaskNodeList = &linkedlist.CircleLinkedList{}
	for scale := 0; scale < domain.SECOND_SCALE; scale++ {
		taskDetailList := make([]MultiLayerTaskDetail, 0, 16)
		taskNode := MultiLayerTaskNode{}
		taskNode.ID = scale + 1
		taskNode.Scale = scale
		taskNode.TaskDetailList = &taskDetailList
		multiLayerTimeWheel.TaskNodeList.AddToTail(&linkedlist.Node{ID: scale, Data: taskNode})
	}
	return multiLayerTimeWheel, nil
}

/**
* 添加只执行一次，不重复执行的任务
 */
func (s *MultiLayerTimeWheel) AddOnceTask(delay int) error {
	return s.AddTask(delay, false)
}

/**
* 添加新任务，可指定是否重复执行
 */
func (s *MultiLayerTimeWheel) AddTask(delay int, repeatFlag bool) error {
	hour, minute, second, scale, err := s.calculateScale(delay)
	if err != nil {
		return err
	}

	taskNode, err := s.findNodeByScale(scale)
	if err != nil {
		return err
	}

	id := strconv.Itoa(taskNode.ID) + "-" + strconv.Itoa(len(*taskNode.TaskDetailList)+1)
	newTask := MultiLayerTaskDetail{}
	newTask.ID = id
	newTask.Scale = scale
	newTask.Hour = hour
	newTask.Minute = minute
	newTask.Second = second
	newTask.currentHour = 0
	newTask.currentMinute = 0
	newTask.currentSecond = 0
	newTask.Delay = delay
	newTask.RepeatFlag = repeatFlag
	newTask.ExecuteFlag = false

	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	newTask.CreateTime = now

	*taskNode.TaskDetailList = append(*taskNode.TaskDetailList, newTask)

	return nil
}

/**
* 删除任务
 */
func (s *MultiLayerTimeWheel) DeleteTask(node *linkedlist.Node) (bool, error) {
	return true, nil
}

/**
* 调度并执行任务
 */
func (s *MultiLayerTimeWheel) ExecuteTask() {
	var wg sync.WaitGroup
	node := s.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for node.Next.ID != s.TaskNodeList.Head.ID {
			taskList := node.Data.(MultiLayerTaskNode).TaskDetailList
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

func (s *MultiLayerTimeWheel) executeSameScaleTask(taskList *[]MultiLayerTaskDetail) {
	for index, task := range *taskList {
		if task.ExecuteFlag {
			continue
		}
		runnable := s.isRunnable(task)
		if runnable {
			fmt.Print("执行分层时间轮任务(id=" + task.ID + ",scale=" + strconv.Itoa(task.Scale) + ",delay=" + strconv.Itoa(task.Delay) + ")，")
			fmt.Println("时间=" + time.Now().Format("2006-01-02 15:04:05"))
			(*taskList)[index].ExecuteFlag = true
			continue
		}
		currentHour := (*taskList)[index].currentHour
		currentMinute := (*taskList)[index].currentMinute
		currentSecond := (*taskList)[index].currentSecond

		addMinuteFlag := false
		addHourFlag := false
		currentSecond += 1
		if currentSecond == 60 {
			currentSecond = 0
			addMinuteFlag = true
		}
		if addMinuteFlag {
			currentMinute += 1
			if currentMinute == 60 {
				currentMinute = 0
				addHourFlag = true
			}
		}
		if addHourFlag {
			if currentHour+1 <= task.Hour {
				currentHour += 1
			}
		}

		if currentHour == task.Hour {
			if currentMinute > task.Minute {
				currentMinute -= 1
				if currentSecond > task.Second {
					currentSecond -= 1
				}
			}
		}

		// fmt.Println("分层时间轮任务(id=" + task.ID + ")最新计时=" + strconv.Itoa(currentHour) + ":" + strconv.Itoa(currentMinute) + ":" + strconv.Itoa(currentSecond))
		(*taskList)[index].currentHour = currentHour
		(*taskList)[index].currentMinute = currentMinute
		(*taskList)[index].currentSecond = currentSecond
	}
}

/**
* 判断任务是否达到运行条件
 */
func (s *MultiLayerTimeWheel) isRunnable(taskDetail MultiLayerTaskDetail) bool {
	if taskDetail.ExecuteFlag {
		return false
	}
	if taskDetail.currentHour != taskDetail.Hour {
		return false
	}
	if taskDetail.currentMinute != taskDetail.Minute {
		return false
	}
	if taskDetail.currentSecond != taskDetail.Second {
		return false
	}
	return true
}

/**
* 找到对应刻度的节点
 */
func (s *MultiLayerTimeWheel) findNodeByScale(scale int) (MultiLayerTaskNode, error) {
	node := s.TaskNodeList.Head.Next
	for node.ID != s.TaskNodeList.Head.ID {
		taskNode := node.Data.(MultiLayerTaskNode)
		if taskNode.Scale == scale {
			return taskNode, nil
		}
		node = node.Next
	}
	return MultiLayerTaskNode{}, errors.New("没有scale=" + strconv.Itoa(scale) + "的节点")
}

/**
* 计算新任务所属的轮次和刻度
 */
func (s *MultiLayerTimeWheel) calculateScale(delay int) (int, int, int, int, error) {
	hour, minute, second, scale := 0, 0, 0, -1
	if delay <= 0 {
		return hour, minute, second, scale, errors.New("延迟时间必须大于0")
	}
	if delay > domain.ONE_DAY_IN_SECOND {
		return hour, minute, second, scale, errors.New("延迟时间不能超过24小时")
	}
	hour = delay / domain.ONE_HOUR_IN_SECOND
	if hour > 0 {
		delay -= hour * 3600
	}
	minute = delay / domain.ONE_MINUTE_IN_SECOND
	if minute > 0 {
		delay -= minute * domain.ONE_MINUTE_IN_SECOND
	}
	second = delay % domain.ONE_MINUTE_IN_SECOND
	scale = second
	return hour, minute, second, scale, nil
}
