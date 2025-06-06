package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"task/internal/metrics"
	"task/internal/worker/processors"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
)

// JobPayload represents the structure of a job payload
type JobPayload struct {
	Type    string          `json:"type"`
	Data    json.RawMessage `json:"data"`
	Timeout int             `json:"timeout"` // timeout in seconds
}

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

				// Log: job started
				rdb.RPush(ctx, "job:"+jobID+":logs", "Started")
				rdb.RPush(ctx, "job:"+jobID+":logs", fmt.Sprintf("Processing job %s: %s (retry %d)", jobID, payload, retryCount))

				log.Printf("Processing job %s: %s (retry %d)", jobID, payload, retryCount)

				// Start timing the job with label
				jobType := "unknown"
				var jobPayload JobPayload
				if err := json.Unmarshal([]byte(payload), &jobPayload); err == nil {
					jobType = jobPayload.Type
				}
				timer := prometheus.NewTimer(metrics.JobDurationByType.WithLabelValues(jobType))
				success := processJob(payload)
				timer.ObserveDuration()

				if !success {
					metrics.JobsFailedTotal.Inc()
					metrics.JobsFailedByType.WithLabelValues(jobType).Inc()
					if retryCount < 3 {
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
						// Log: retry
						rdb.RPush(ctx, "job:"+jobID+":logs", fmt.Sprintf("Retry %d", retryCount+1))
					} else {
						// Log: failed
						rdb.RPush(ctx, "job:"+jobID+":logs", "Failed")
					}
				} else {
					metrics.JobsProcessedTotal.Inc()
					metrics.JobsProcessedByType.WithLabelValues(jobType).Inc()
					rdb.XAck(ctx, "jobs", "workers", message.ID)
					// Log: success
					rdb.RPush(ctx, "job:"+jobID+":logs", "Success")
				}

				status := "success"
				if !success {
					status = "failed"
				}

				rdb.HSet(ctx, "job:"+jobID, map[string]interface{}{
					"status":     status,
					"retry":      retryCount,
					"updated_at": time.Now().Format(time.RFC3339),
				})
			}
		}
	}
}

func processJob(payload string) bool {
	var jobPayload JobPayload
	if err := json.Unmarshal([]byte(payload), &jobPayload); err != nil {
		log.Printf("Error parsing job payload: %v", err)
		return false
	}

	// Create a context with timeout if specified
	ctx := context.Background()
	if jobPayload.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(jobPayload.Timeout)*time.Second)
		defer cancel()
	}

	// Process different job types
	switch jobPayload.Type {
	case "email":
		return processors.ProcessEmailJob(ctx, jobPayload.Data)
	case "notification":
		return processNotificationJob(ctx, jobPayload.Data)
	case "cleanup":
		return processCleanupJob(ctx, jobPayload.Data)
	default:
		log.Printf("Unknown job type: %s", jobPayload.Type)
		return false
	}
}

func processEmailJob(ctx context.Context, data json.RawMessage) bool {
	// TODO: Implement email sending logic
	log.Printf("Processing email job with data: %s", string(data))
	return true
}

func processNotificationJob(ctx context.Context, data json.RawMessage) bool {
	// TODO: Implement notification logic
	log.Printf("Processing notification job with data: %s", string(data))
	return true
}

func processCleanupJob(ctx context.Context, data json.RawMessage) bool {
	// TODO: Implement cleanup logic
	log.Printf("Processing cleanup job with data: %s", string(data))
	return true
}
