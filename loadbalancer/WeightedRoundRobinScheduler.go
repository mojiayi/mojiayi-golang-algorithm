package loadbalancer

type WeightedRoundRobinScheduler struct{}

/**
 * 加权轮询，不管有多少个服务节点，权重值分别是多少，每运行@{code sum(server1.weight+server2.weight+...+serverN.weight)/2}个轮次，各服务节点的动态权重恢复默认值，从头开始分配
 *
 */
func (s WeightedRoundRobinScheduler) Choose(serverList *[]ServerInfo) ServerInfo {
	var needInitCurrentWeight = true
	for _, value := range *serverList {
		if value.CurrentWeight != 0 {
			needInitCurrentWeight = false
			break
		}
	}
	if needInitCurrentWeight {
		for index, value := range *serverList {
			(*serverList)[index].CurrentWeight = value.Weight
		}
	}
	var chosenServer = (*serverList)[0]
	var chosenIndex = 0
	for index := 0; index < len(*serverList); index++ {
		if index == 0 {
			continue
		} else {
			if chosenServer.CurrentWeight < (*serverList)[index].CurrentWeight {
				chosenServer = (*serverList)[index]
				chosenIndex = index
			}
		}
	}

	(*serverList)[chosenIndex].CurrentWeight = chosenServer.CurrentWeight - 1
	return chosenServer
}
