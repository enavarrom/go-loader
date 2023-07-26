package writer

import (
	"fileLoader/config"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetItemWriter() ItemWriter {
	itemWriter, err := redisStreamItemWriter()
	if err != nil {
		log.Fatal(err)
	}
	return itemWriter
}

func redisStreamItemWriter() (*RedisStreamItemWriter, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		fmt.Println("Error Redis Port:", err)
		return &RedisStreamItemWriter{}, err
	}
	return NewRedisStreamItemWriter(redisHost, redisPort, config.GetInstance().StreamName), nil
}
