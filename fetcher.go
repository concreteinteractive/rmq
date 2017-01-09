package rmq

import (
	"gopkg.in/redis.v5"
)

type Fetcher interface {
	Fetch(redisClient *redis.Client, queue *redisQueue) string
}
