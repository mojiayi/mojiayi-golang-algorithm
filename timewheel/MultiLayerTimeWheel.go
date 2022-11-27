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
	* 所在层级
	 */
	Layer domain.LayerName
	/**
	* 上次变化刻度的时间，单位为秒
	 */
	LatestUpdateTime int
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
* 计时钟表
 */
type TimingClock struct {
	HourPointer   *MultiLayerTimeWheel
	MinutePointer *MultiLayerTimeWheel
	SecondPointer *MultiLayerTimeWheel
}

/**
* 创建一个分层时间轮
 */
func (s *TimingClock) New() (*TimingClock, error) {
	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	timingClock := new(TimingClock)

	timingClock.HourPointer, _ = s.newPointer(domain.HOUR_SCALE, now)
	timingClock.MinutePointer, _ = s.newPointer(domain.MINITE_SCALE, now)
	timingClock.SecondPointer, _ = s.newPointer(domain.SECOND_SCALE, now)

	return timingClock, nil
}

func (s *TimingClock) newPointer(maxScale int, now int) (*MultiLayerTimeWheel, error) {
	timeWheel := new(MultiLayerTimeWheel)
	timeWheel.MaxScale = maxScale
	timeWheel.CurrentScale = 0
	timeWheel.StartupTime = now
	timeWheel.CurrentTime = now
	timeWheel.TaskNodeList = &linkedlist.CircleLinkedList{}
	for scale := 0; scale < maxScale; scale++ {
		taskDetailList := make([]MultiLayerTaskDetail, 0, 16)
		taskNode := MultiLayerTaskNode{}
		taskNode.ID = scale
		taskNode.Scale = scale
		taskNode.TaskDetailList = &taskDetailList
		timeWheel.TaskNodeList.AddToTail(&linkedlist.Node{ID: scale, Data: taskNode})
	}
	return timeWheel, nil
}

/**
* 添加只执行一次，不重复执行的任务
 */
func (s *TimingClock) AddOnceTask(delay int) error {
	return s.AddTask(delay, false)
}

/**
* 添加新任务，可指定是否重复执行
 */
func (s *TimingClock) AddTask(delay int, repeatFlag bool) error {
	hour, minute, second, err := s.calculateScale(delay)
	if err != nil {
		return err
	}

	taskNode, layer, err := s.findNodeByScale(hour, minute, second)
	if err != nil {
		return err
	}

	now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
	id := strconv.Itoa(taskNode.ID) + "-" + strconv.Itoa(len(*taskNode.TaskDetailList)+1)
	newTask := MultiLayerTaskDetail{}
	newTask.ID = id
	newTask.Hour = hour
	newTask.Minute = minute
	newTask.Second = second
	newTask.currentHour = 0
	newTask.currentMinute = 0
	newTask.currentSecond = 0
	newTask.Delay = delay
	newTask.RepeatFlag = repeatFlag
	newTask.ExecuteFlag = false
	newTask.Layer = layer
	newTask.LatestUpdateTime = now
	newTask.CreateTime = now

	*taskNode.TaskDetailList = append(*taskNode.TaskDetailList, newTask)

	return nil
}

/**
* 调度并执行任务
 */
func (s *TimingClock) ExecuteTask() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		s.driveSecondPointer()
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		s.driveMinutePointer()
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		s.driveHourPointer()
		defer wg.Done()
	}()

	wg.Wait()
}

func (s *TimingClock) driveSecondPointer() {
	var wg sync.WaitGroup
	node := s.SecondPointer.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for node.Next.ID != s.SecondPointer.TaskNodeList.Head.ID {
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
		node = s.SecondPointer.TaskNodeList.Head
	}
}

func (s *TimingClock) driveMinutePointer() {
	var wg sync.WaitGroup
	node := s.MinutePointer.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for node.Next.ID != s.MinutePointer.TaskNodeList.Head.ID {
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
		node = s.MinutePointer.TaskNodeList.Head
	}
}

func (s *TimingClock) driveHourPointer() {
	var wg sync.WaitGroup
	node := s.HourPointer.TaskNodeList.Head

	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		for node.Next.ID != s.HourPointer.TaskNodeList.Head.ID {
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
		node = s.HourPointer.TaskNodeList.Head
	}
}

func (s *TimingClock) executeSameScaleTask(taskList *[]MultiLayerTaskDetail) {
	var removeIndexes = make([]int, len(*taskList))
	for index, task := range *taskList {
		removeIndexes[index] = -1
		if task.ExecuteFlag {
			continue
		}
		runnable := s.isRunnable(task)
		if runnable {
			fmt.Print("执行分层时间轮任务(id=" + task.ID + ",delay=" + strconv.Itoa(task.Delay) + ")，")
			fmt.Println("时间=" + time.Now().Format("2006-01-02 15:04:05"))
			(*taskList)[index].ExecuteFlag = true
			continue
		}

		if (*taskList)[index].Layer == domain.LAYER_SECOND {
			currentSecond := (*taskList)[index].currentSecond + 1
			(*taskList)[index].currentSecond = currentSecond
			(*taskList)[index].LatestUpdateTime += 1
			if currentSecond == (*taskList)[index].Second {
				if (*taskList)[index].Minute > 0 {
					minutePointerTaskNode, layer, err := s.findNodeByScale(0, (*taskList)[index].Minute, 0)
					if err != nil {
						return
					} else {
						(*taskList)[index].Layer = layer
						*minutePointerTaskNode.TaskDetailList = append(*minutePointerTaskNode.TaskDetailList, (*taskList)[index])
						removeIndexes[index] = index
					}
				} else if (*taskList)[index].Hour > 0 {
					hourPointerTaskNode, layer, err := s.findNodeByScale((*taskList)[index].Hour, 0, 0)
					if err != nil {
						return
					} else {
						(*taskList)[index].Layer = layer
						*hourPointerTaskNode.TaskDetailList = append(*hourPointerTaskNode.TaskDetailList, (*taskList)[index])
						removeIndexes[index] = index
					}
				}
			}
		}
		if (*taskList)[index].Layer == domain.LAYER_MINUTE {
			now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
			if now-(*taskList)[index].LatestUpdateTime >= domain.ONE_MINUTE_IN_SECOND {
				(*taskList)[index].LatestUpdateTime += domain.ONE_MINUTE_IN_SECOND

				currentMinute := (*taskList)[index].currentMinute + 1
				(*taskList)[index].currentMinute = currentMinute
				if currentMinute == (*taskList)[index].Minute {
					if (*taskList)[index].Hour > 0 {
						hourPointerTaskNode, layer, err := s.findNodeByScale((*taskList)[index].Hour, 0, 0)
						if err != nil {
							return
						} else {
							(*taskList)[index].Layer = layer
							*hourPointerTaskNode.TaskDetailList = append(*hourPointerTaskNode.TaskDetailList, (*taskList)[index])
							removeIndexes[index] = index
						}
					}
				}
			}
		}
		if (*taskList)[index].Layer == domain.LAYER_HOUR {
			now := int(time.Now().UnixMilli() / int64(domain.ONE_THOUSAND))
			if now-(*taskList)[index].LatestUpdateTime >= domain.ONE_HOUR_IN_SECOND {
				currentHour := (*taskList)[index].currentHour + 1
				(*taskList)[index].currentHour = currentHour
			}
		}
	}

	for _, index := range removeIndexes {
		if index < 0 {
			continue
		}
		*taskList = append((*taskList)[:index], (*taskList)[index+1:]...)
	}
}

/**
* 判断任务是否达到运行条件
 */
func (s *TimingClock) isRunnable(taskDetail MultiLayerTaskDetail) bool {
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
func (s *TimingClock) findNodeByScale(hour int, minute int, second int) (MultiLayerTaskNode, domain.LayerName, error) {
	if second > 0 {
		node := s.SecondPointer.TaskNodeList.Head.Next
		for node.ID != s.SecondPointer.TaskNodeList.Head.ID {
			taskNode := node.Data.(MultiLayerTaskNode)
			if taskNode.Scale == second {
				return taskNode, domain.LAYER_SECOND, nil
			}
			node = node.Next
		}
	}
	if minute > 0 {
		node := s.MinutePointer.TaskNodeList.Head.Next
		for node.ID != s.MinutePointer.TaskNodeList.Head.ID {
			taskNode := node.Data.(MultiLayerTaskNode)
			if taskNode.Scale == minute {
				return taskNode, domain.LAYER_MINUTE, nil
			}
			node = node.Next
		}
	}
	if hour > 0 {
		node := s.HourPointer.TaskNodeList.Head.Next
		for node.ID != s.HourPointer.TaskNodeList.Head.ID {
			taskNode := node.Data.(MultiLayerTaskNode)
			if taskNode.Scale == hour {
				return taskNode, domain.LAYER_HOUR, nil
			}
			node = node.Next
		}
	}
	return MultiLayerTaskNode{}, domain.UNKNOWN, errors.New("没有刻度=" + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second) + "的节点")
}

/**
* 计算新任务所属的时分秒
 */
func (s *TimingClock) calculateScale(delay int) (int, int, int, error) {
	hour, minute, second := -1, -1, -1
	if delay <= 0 {
		return hour, minute, second, errors.New("延迟时间必须大于0")
	}
	if delay > domain.ONE_DAY_IN_SECOND {
		return hour, minute, second, errors.New("延迟时间不能超过24小时")
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
	return hour, minute, second, nil
}
