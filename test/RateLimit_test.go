package algorithm_test

import (
	"fmt"
	"math/rand"
	"mojiayi-golang-algorithm/ratelimiter"
	"strconv"
	"testing"
	"time"
)

func TestLeakBucket(t *testing.T) {
	var instance = ratelimiter.LeakBucket{LastAccessTimeMap: map[string]int64{}, Qps: 5, Interval: 1000}

	for i := 0; i < 10; i++ {
		isAcquire, difference := instance.TryAcquire(123, "abc")
		fmt.Printf("now is %d,result=%s,difference=%d", time.Now().UnixMilli(), strconv.FormatBool(isAcquire), difference)
		fmt.Println()

		time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
	}
}

func TestTokenBucket(t *testing.T) {
	var instance = ratelimiter.TokenBucket{LastAccessTimeMap: map[string]int64{}, RemainTokenMap: map[string]int{}, Qps: 5, Interval: 1000}

	for i := 0; i < 10; i++ {
		isAcquire, difference, remainToken := instance.TryAcquire(123, "abc")
		fmt.Printf("now is %d,result=%s,difference=%d,remainToken=%d", time.Now().UnixMilli(), strconv.FormatBool(isAcquire), difference, remainToken)
		fmt.Println()

		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}
