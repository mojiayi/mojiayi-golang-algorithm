package algorithm_test

import (
	"fmt"
	"mojiayi-golang-algorithm/domain"
	"mojiayi-golang-algorithm/loadbalancer"
	"strconv"
	"sync"
	"testing"
)

var (
	serverSize = 5
	hosts      = make([]domain.ServerInfo, serverSize)

	weights = []int{2, 6, 8, 7}
	hosts2  = make([]domain.ServerInfo, len(weights))
)

func init() {
	for index := 0; index < serverSize; index++ {
		var instance domain.ServerInfo
		instance.ServiceId = index
		instance.Weight = index + 1
		instance.ServiceName = "testService" + strconv.Itoa(index)
		hosts[index] = instance
	}

	for index := 0; index < len(weights); index++ {
		var instance domain.ServerInfo
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
	var wg sync.WaitGroup
	var done = make(chan bool, 1)
	var serverScheduler = loadbalancer.RoundRobinScheduler{}
	done <- true
	var cycle = 100
	var chosenArr = make([]int, cycle)
	for count := 0; count < cycle; count++ {
		wg.Add(1)
		go func(index int, chosenArr []int) {
			<-done
			defer wg.Done()
			var chosenInstance = serverScheduler.Choose(&hosts)

			chosenArr[index] = chosenInstance.Weight
			done <- true

		}(count, chosenArr)
	}
	wg.Wait()

	for index, weight := range chosenArr {
		fmt.Printf("chosen node index=%v,weight=%v\n", index, weight)
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
