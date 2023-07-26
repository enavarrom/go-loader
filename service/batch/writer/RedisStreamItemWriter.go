package writer

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisStreamItemWriter struct {
	client     *redis.Client
	context    context.Context
	streamName string
}

func NewRedisStreamItemWriter(redisHost string, redisPort int, streamName string) *RedisStreamItemWriter {
	return &RedisStreamItemWriter{
		client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", redisHost, redisPort),
		}),
		context:    context.Background(),
		streamName: streamName}
}

func (c *RedisStreamItemWriter) Write(items []interface{}) {
	for _, item := range items {
		_, err := c.client.XAdd(c.context, &redis.XAddArgs{
			Stream: c.streamName,
			Values: map[string]interface{}{"message": item},
		}).Result()
		if err != nil {
			log.Fatal("Error al agregar mensaje:", err)
		}
	}
}
