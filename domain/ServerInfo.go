package domain

/**
 * 服务节点基本信息，仅用于试验各种负载均衡分配算法。
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
	 * 专门用于按权重轮询的动态权重值，在分配过程中变化
	 */
	CurrentWeight int
}
