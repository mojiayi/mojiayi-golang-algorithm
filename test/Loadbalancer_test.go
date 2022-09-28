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
		instance.Weight = weights[index]
		instance.CurrentWeight = weights[index]
		instance.ServiceName = "testService" + strconv.Itoa(index)
		hosts2[index] = instance
	}
}

func TestWeightedRandomScheduler(t *testing.T) {
	var wg sync.WaitGroup
	var done = make(chan bool, 1)
	var serverScheduler = loadbalancer.WeightedRandomScheduler{}
	done <- true
	var cycle = 20
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

func TestRoundRobinScheduler(t *testing.T) {
	var wg sync.WaitGroup
	var done = make(chan bool, 1)
	var serverScheduler = loadbalancer.RoundRobinScheduler{}
	done <- true
	var cycle = 20
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
	var wg sync.WaitGroup
	var done = make(chan bool, 1)
	var serverScheduler = loadbalancer.WeightedRoundRobinScheduler{}
	done <- true
	var cycle = 20
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

func TestSmoothWeightedRoundRobinScheduler(t *testing.T) {
	var wg sync.WaitGroup
	var done = make(chan bool, 1)
	var serverScheduler = loadbalancer.SmoothWeightedRoundRobinScheduler{}
	done <- true
	var cycle = 20
	var chosenArr = make([]int, cycle)
	for count := 0; count < cycle; count++ {
		wg.Add(1)
		go func(count int, chosenArr []int) {
			<-done
			defer wg.Done()
			var chosenInstance = serverScheduler.Choose(&hosts2)

			chosenArr[count] = chosenInstance.Weight
			done <- true

		}(count, chosenArr)
	}
	wg.Wait()

	for index, weight := range chosenArr {
		fmt.Printf("chosen node index=%v,weight=%v\n", index, weight)
	}
}
