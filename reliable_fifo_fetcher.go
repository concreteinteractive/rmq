package rmq

import (
	"github.com/go-redis/redis"
)

type ReliableFifoFetcher struct {
}

func (f *ReliableFifoFetcher) Fetch(redisClient *redis.Client, queue *redisQueue) string {
	result := redisClient.RPopLPush(queue.readyKey, queue.unackedKey)
	if redisErrIsNil(result) {
		return ""
	}

	return result.Val()
}
