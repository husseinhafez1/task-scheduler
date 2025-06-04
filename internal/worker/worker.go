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
	_ = rdb.XGroupCreateMkStream(ctx, "jobs", "workers", "$")
	for {
		entries, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "workers",
			Consumer: "worker-1",
			Streams:  []string{"jobs", ">"},
			Block:    0,
			Count:    1,
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

				log.Printf("Processing job %s: %s (retry %d)", jobID, payload, retryCount)
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
				if success {
					rdb.XAck(ctx, "jobs", "workers", message.ID)
				}
				rdb.HSet(ctx, "job:"+jobID, map[string]interface{}{
					"status":     "success", // or "failed"
					"retry":      retryCount,
					"updated_at": time.Now().Format(time.RFC3339),
				})
			}
		}
	}

}

func processJob(payload string) bool {
	log.Printf("Processing payload: %s", payload)
	return true
}
