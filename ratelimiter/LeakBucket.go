package ratelimiter

import (
	"strconv"
	"sync"
	"time"
)

type LeakBucket struct{}

var lastAccessTimeMap = make(map[string]int64)
var maxTimes = 5

func (c LeakBucket) TryAcquire(userId int64, uri string) bool {
	mutex := &sync.Mutex{}

	locked := mutex.TryLock()
	isAcquire := false
	if !locked {
		return isAcquire
	}
	var key = strconv.Itoa(int(userId)) + "-" + uri
	value, ok := lastAccessTimeMap[key]
	nowMilli := time.Now().UnixMilli()
	if !ok || value <= 0 {
		lastAccessTimeMap[key] = nowMilli
		isAcquire = true
	} else {
		// 最简单的示例，每秒5次
		difference := nowMilli - value
		ok = difference-int64(maxTimes*200) >= 0
		if ok {
			lastAccessTimeMap[key] = nowMilli
			isAcquire = true
		}
	}

	mutex.Unlock()
	return isAcquire
}
