package sorting_test

import (
	"fmt"
	"mojiayi-golang-algorithm/loadbalancer"
	"strconv"
	"testing"
)

var (
	serverSize = 5
	hosts      = make([]loadbalancer.ServerInfo, serverSize)

	weights = []int{2, 6, 8, 7}
	hosts2  = make([]loadbalancer.ServerInfo, len(weights))
)

func init() {
	for index := 0; index < serverSize; index++ {
		var instance loadbalancer.ServerInfo
		instance.ServiceId = index
		instance.Weight = index + 1
		instance.ServiceName = "testService" + strconv.Itoa(index)
		hosts[index] = instance
	}

	for index := 0; index < len(weights); index++ {
		var instance loadbalancer.ServerInfo
		instance.ServiceId = index
		instance.Weight = index + 1
		instance.ServiceName = "testService" + strconv.Itoa(index)
		hosts2[index] = instance
	}
}

func TestWeightedRandomScheduler(t *testing.T) {
	var serverScheduler loadbalancer.WeightedRandomScheduler
	var cycle = 10
	var matchRecord = choose(serverScheduler, cycle)
	fmt.Println(matchRecord)
}

func TestRoundRobinScheduler(t *testing.T) {
	var serverScheduler loadbalancer.RoundRobinScheduler
	var cycle = 10
	for count := 0; count < cycle; count++ {
		var chosenInstance = serverScheduler.Choose(&hosts)
		var expected = (count % serverSize) + 1
		if chosenInstance.Weight != expected {
			t.Errorf("chosenInstance[%v] excepted:%v,actual:%v", chosenInstance, expected, chosenInstance.Weight)
		}
	}
}

func TestWeightedRoundRobinScheduler(t *testing.T) {
	var serverScheduler loadbalancer.WeightedRoundRobinScheduler
	var cycle = 100
	for count := 0; count < cycle; count++ {
		var chosenInstance = serverScheduler.Choose(&hosts)
		fmt.Println(chosenInstance.ServiceName)
	}
}

func TestSmoothWeightedRoundRobinScheduler(t *testing.T) {
	var serverScheduler loadbalancer.SmoothWeightedRoundRobinScheduler
	var cycle = 100
	for count := 0; count < cycle; count++ {
		var chosenInstance = serverScheduler.Choose(&hosts)
		fmt.Println(chosenInstance.ServiceName)
	}
}

func choose(serverScheduler loadbalancer.IServerScheduler, cycle int) []string {
	var matchRecord = make([]string, cycle)
	for count := 0; count < cycle; count++ {
		var chosenInstance = serverScheduler.Choose(&hosts)
		matchRecord[count] = chosenInstance.ServiceName
	}
	return matchRecord
}
