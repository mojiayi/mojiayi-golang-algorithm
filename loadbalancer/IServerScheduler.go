package loadbalancer

/**
 * 按具体实现类中的方式选择服务节点
 */
type IServerScheduler interface {
	/**
	 * 返回一个服务节点
	 *
	 * @param serverList 可被分配的服务节点列表
	 * @return 传入的服务节点列表为空时，返回空，否则，一定会返回一个节点
	 */
	Choose(serverList *[]ServerInfo) ServerInfo
}

/**
 * 服务节点基本信息，仅用于试验各种负载均衡分配算法，所以信息很少。
 *
 */
type ServerInfo struct {
	/**
	 * 服务id
	 */
	ServiceId int
	/**
	 * 服务节点名
	 */
	ServiceName string
	/**
	 * 服务节点固定权重值，在分配过程中不变
	 */
	Weight int
	/**
	 * 专门用于按权重轮询的动态权重值，在{@code ChooseByRoundRobinWeight}的分配过程中变化
	 */
	CurrentWeight int
}
