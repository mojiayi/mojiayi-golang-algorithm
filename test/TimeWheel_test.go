package algorithm_test

import (
	"mojiayi-golang-algorithm/timewheel"
	"testing"
)

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

	instance.AddOnceTask(62)
	instance.AddOnceTask(9)
	instance.AddOnceTask(2)
	instance.AddOnceTask(69)

	instance.ExecuteTask()
}

func TestMultiLayerTimeWheel(t *testing.T) {
	multiLayerTimeWheel := timewheel.MultiLayerTimeWheel{}
	instance, err := multiLayerTimeWheel.New()
	if err != nil {
		t.Errorf("构建分层时间轮失败%v", err)
	}

	instance.AddOnceTask(62)
	instance.AddOnceTask(9)
	instance.AddOnceTask(2)
	instance.AddOnceTask(69)

	instance.ExecuteTask()
}
