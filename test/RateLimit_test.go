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
	var instance ratelimiter.LeakBucket
	for i := 0; i < 10; i++ {
		fmt.Printf("now is %d,result=%s", time.Now().UnixMilli(), strconv.FormatBool(instance.TryAcquire(123, "abc")))
		fmt.Println()

		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}

}
