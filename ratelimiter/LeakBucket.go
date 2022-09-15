package ratelimiter

import (
	"strconv"
	"sync"
	"time"
)

/**
* 漏桶算法的简单实现
 */
type LeakBucket struct {
	LastAccessTimeMap map[string]int64
	Qps               int
	Interval          int64
}

func (c LeakBucket) TryAcquire(userId int64, uri string) bool {
	mutex := &sync.Mutex{}

	// 尝试获取本地锁，如果获取失败，直接返回
	locked := mutex.TryLock()
	isAcquire := false
	if !locked {
		return isAcquire
	}
	var key = strconv.Itoa(int(userId)) + "-" + uri
	lastAccessTime, ok := c.LastAccessTimeMap[key]
	nowMilli := time.Now().UnixMilli()
	if !ok {
		// 指定的key还没有访问记录，作为初次访问对待
		c.LastAccessTimeMap[key] = nowMilli
		isAcquire = true
	} else {
		// 计算上次访问到现在是否有足够长的间隔
		difference := nowMilli - lastAccessTime
		ok = difference-(c.Interval/(int64)(c.Qps)) >= 0
		if ok {
			c.LastAccessTimeMap[key] = nowMilli
			isAcquire = true
		}
	}

	mutex.Unlock()
	return isAcquire
}
