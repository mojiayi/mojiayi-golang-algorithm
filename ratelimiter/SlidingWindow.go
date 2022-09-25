package ratelimiter

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

/**
* 滑动窗口算法的简单实现
 */
type SlidingWindow struct {
	Capacity        uint64
	leftPointIndex  uint64
	rightPointIndex uint64
}

func (c *SlidingWindow) TryAcquire(wg *sync.WaitGroup) (bool, uint64) {
	defer wg.Done()

	// 尝试获取本地锁，如果获取失败，直接返回
	// c.Lock()
	var isAcquire bool
	var difference = c.rightPointIndex - c.leftPointIndex
	if difference >= c.Capacity {
		isAcquire = false
	} else {
		isAcquire = true
		atomic.AddUint64(&c.rightPointIndex, 1)
		time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
		atomic.AddUint64(&c.leftPointIndex, 1)
	}
	// c.Unlock()

	fmt.Printf("now is %d,result=%s,difference=%d", time.Now().UnixMilli(), strconv.FormatBool(isAcquire), difference)
	fmt.Println()
	return isAcquire, difference
}
