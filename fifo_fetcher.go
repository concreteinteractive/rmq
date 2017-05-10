package rmq

import (
	"github.com/go-redis/redis"
)

type FifoFetcher struct {
}

func (f *FifoFetcher) Fetch(redisClient *redis.Client, queue *redisQueue) string {
	result := redisClient.RPop(queue.readyKey)
	if redisErrIsNil(result) {
		return ""
	}
	return result.Val()
}
