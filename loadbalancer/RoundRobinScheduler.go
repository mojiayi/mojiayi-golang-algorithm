package loadbalancer

type RoundRobinScheduler struct{}

var currentServerIndex = -1

/**
* 普通轮询
 */
func (s RoundRobinScheduler) Choose(serverList *[]ServerInfo) ServerInfo {
	var nextServerIndex = currentServerIndex
	if nextServerIndex < 0 {
		nextServerIndex = 0
		currentServerIndex = 0
	} else {
		currentServerIndex++
		if currentServerIndex == len(*serverList) {
			nextServerIndex = 0
			currentServerIndex = 0
		} else {
			nextServerIndex = currentServerIndex
		}
	}
	return (*serverList)[nextServerIndex]
}
