package rmq

import (
	"github.com/go-redis/redis"
)

type Fetcher interface {
	Fetch(redisClient *redis.Client, queue *redisQueue) string
}
