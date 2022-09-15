package ratelimiter

type IRateLimiter interface {
	TryAcquire(userId int64, uri string) bool
}
