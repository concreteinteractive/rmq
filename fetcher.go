package rmq

import (
	"gopkg.in/redis.v3"
)

type Fetcher interface {
	Fetch(redisClient *redis.Client, queue *redisQueue) string
}
