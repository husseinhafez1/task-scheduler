package worker

import (
	"context"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func StartWorker(rdb *redis.Client) {
	ctx := context.Background()
	for {
		entries, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"jobs", "0"},
			Block:   0,
		}).Result()

		if err != nil {
			log.Printf("Error reading from Redis: %v", err)
			continue
		}

		for _, stream := range entries {
			for _, message := range stream.Messages {
				jobID := message.Values["job_id"].(string)
				payload := message.Values["payload"].(string)
				retryCount, _ := strconv.Atoi(message.Values["retry"].(string))

				log.Printf("Processing job %s: %s", jobID, payload, retryCount)
				success := processJob(payload)

				if !success && retryCount < 3 {
					backoff := time.Duration(math.Pow(2, float64(retryCount))) * time.Second
					time.Sleep(backoff)

					rdb.XAdd(ctx, &redis.XAddArgs{
						Stream: "jobs",
						Values: map[string]interface{}{
							"job_id":  jobID,
							"payload": payload,
							"retry":   retryCount + 1,
						},
					})
				}
			}
		}
	}

}

func processJob(payload string) bool {
	log.Printf("Processing payload: %s", payload)
	return true
}
