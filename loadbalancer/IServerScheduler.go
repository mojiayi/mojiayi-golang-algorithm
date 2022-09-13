package loadbalancer

import "mojiayi-golang-algorithm/domain"

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
	Choose(serverList *[]domain.ServerInfo) domain.ServerInfo
}
