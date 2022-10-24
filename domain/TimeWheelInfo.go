package domain

import "mojiayi-golang-algorithm/linkedlist"

var (
	MAX_SCALE            int = 60
	ONE_THOUSAND             = 1000
	ONE_DAY_IN_SECOND        = 86400
	ONE_HOUR_IN_SECOND       = 3600
	ONE_MINUTE_IN_SECOND     = 60
	ONE_DAY_IN_HOUR          = 24
)

type LayerName int

const (
	LAYER_HOUR   = 1
	LAYER_MINUTE = 2
	LAYER_SECOND = 3
)

type MojiayiTimeWheel struct {
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
	* 是否已经执行
	 */
	ExecuteFlag bool
	/**
	* 是否重复执行
	 */
	RepeatFlag bool
	/**
	* 创建时间，单位为秒
	 */
	CreateTime int
}
