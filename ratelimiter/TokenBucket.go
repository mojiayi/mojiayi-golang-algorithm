package ratelimiter

import (
	"strconv"
	"sync"
	"time"
)

/**
* 令牌桶算法的简单实现
 */
type TokenBucket struct {
	LastAccessTimeMap map[string]int64
	RemainTokenMap    map[string]int
	Qps               int
	Interval          int64
}

func (c *TokenBucket) TryAcquire(userId int64, uri string) (bool, int64, int) {
	mutex := &sync.Mutex{}

	// 尝试获取本地锁，如果获取失败，直接返回
	locked := mutex.TryLock()
	var isAcquire bool
	var difference int64
	var remainToken int
	if !locked {
		return isAcquire, int64(difference), remainToken
	}
	var key = strconv.Itoa(int(userId)) + "-" + uri
	lastAccessTime, ok := c.LastAccessTimeMap[key]
	nowMilli := time.Now().UnixMilli()
	if !ok {
		// 指定的key还没有访问记录，作为初次访问对待
		c.LastAccessTimeMap[key] = nowMilli
		c.RemainTokenMap[key] = c.Qps - 1
		isAcquire = true
	} else {
		remainToken, ok = c.RemainTokenMap[key]
		if !ok {
			// 指定的key还没有访问记录，作为初次访问对待
			c.LastAccessTimeMap[key] = nowMilli
			c.RemainTokenMap[key] = c.Qps - 1
			isAcquire = true
		} else {
			difference = time.Since(time.UnixMilli(lastAccessTime)).Milliseconds()
			if difference >= c.Interval {
				// 上次访问的时间间隔已经超过指定频率，作为初次访问对待
				c.LastAccessTimeMap[key] = nowMilli
				c.RemainTokenMap[key] = c.Qps - 1
				isAcquire = true
			} else {
				// 计算上次访问到现在剩余+新产生的可用次数
				remainToken = int(float64(difference)/float64(c.Interval)*float64(c.Qps)) + remainToken
				if remainToken > 0 {
					// 可用次数大于0时，更新访问记录
					c.LastAccessTimeMap[key] = nowMilli
					c.RemainTokenMap[key] = remainToken - 1
					if c.RemainTokenMap[key] >= c.Qps {
						c.RemainTokenMap[key] = c.Qps - 1
					}
					isAcquire = true
				}
			}
		}
	}

	mutex.Unlock()
	return isAcquire, int64(difference), remainToken
}
