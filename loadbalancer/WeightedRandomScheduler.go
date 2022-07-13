package loadbalancer

import "math/rand"

type WeightedRandomScheduler struct{}

/**
* 加权随机，实现思路来自Ribbon源码。
 */
func (s WeightedRandomScheduler) Choose(serverList *[]ServerInfo) ServerInfo {
	// 把所有服务节点的权重值相加，用于后续的计算
	var weightSum float64
	for _, item := range *serverList {
		weightSum += float64(item.Weight)
	}
	var exactWeights = make([]float64, len(*serverList))

	for index, item := range *serverList {
		// 计算得出每个节点的权重值在总权重值之和中占多少比例
		exactWeights[index] = float64(item.Weight) / weightSum
	}
	var weights = make([]float64, len(*serverList))
	for i := 0; i < len(exactWeights); i++ {
		// 从第2个服务节点起，后一个服务节点本身的权重值+前一个，形成类似等差数列的一个权重值数组
		// 最后一个节点的权重值应该是1，或者小于1且无限接近1的数值
		weights[i] = exactWeights[i]
		if i > 0 {
			weights[i] = weights[i] + weights[i-1]
		}
	}
	// 生成一个0~1，不包含1的伪随机数
	var random = rand.Float64()
	selectedIndex := -1
	for i := 0; i < len(weights); i++ {
		if weights[i] <= random && weights[i+1] >= random {
			selectedIndex = i
		}
	}

	if selectedIndex < 0 {
		selectedIndex = -selectedIndex - 1
	} else {
		return (*serverList)[selectedIndex]
	}

	if selectedIndex < len(weights) {
		return (*serverList)[selectedIndex]
	}

	// 这行代码一般不会被运行到，除非前面的计算过程中发生了浮点数精度问题
	// 为防止这种极小概率的事件发生，默认返回最后一个服务节点
	return (*serverList)[len(*serverList)-1]
}
