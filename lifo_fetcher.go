package rmq

import (
	"gopkg.in/redis.v5"
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
