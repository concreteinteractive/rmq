package rmq

import (
	"gopkg.in/redis.v5"
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
