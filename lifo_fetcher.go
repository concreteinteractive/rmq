package rmq

import (
	"github.com/go-redis/redis"
)

type LifoFetcher struct {
}

func (f *LifoFetcher) Fetch(redisClient *redis.Client, queue *redisQueue) string {
	result := redisClient.LPop(queue.readyKey)
	if redisErrIsNil(result) {
		return ""
	}
	return result.Val()
}
