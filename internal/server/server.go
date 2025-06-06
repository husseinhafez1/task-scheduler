package server

import (
	"context"
	"log"

	"task/internal/metrics"
	pb "task/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	Rdb *redis.Client
}

func NewServer(rdb *redis.Client) *TaskServer {
	return &TaskServer{Rdb: rdb}
}

func (s *TaskServer) GetJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	metrics.JobsProcessedByType.WithLabelValues("GetJobStatus").Inc()
	if req.JobId == "" {
		metrics.JobsFailedByType.WithLabelValues("GetJobStatus").Inc()
		log.Printf("GetJobStatus: missing job_id")
		return nil, status.Error(codes.InvalidArgument, "job_id is required")
	}
	status, err := s.Rdb.HGet(ctx, "job:"+req.JobId, "status").Result()
	if err == redis.Nil {
		return &pb.JobStatusResponse{Status: "not found"}, nil
	} else if err != nil {
		metrics.JobsFailedByType.WithLabelValues("GetJobStatus").Inc()
		log.Printf("GetJobStatus: redis error: %v", err)
		return nil, err
	}
	log.Printf("GetJobStatus: job_id=%s status=%s", req.JobId, status)
	return &pb.JobStatusResponse{Status: status}, nil
}

func (s *TaskServer) SubmitJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	metrics.JobsProcessedByType.WithLabelValues("SubmitJob").Inc()
	if req.JobId == "" || req.Payload == "" {
		metrics.JobsFailedByType.WithLabelValues("SubmitJob").Inc()
		log.Printf("SubmitJob: missing job_id or payload")
		return nil, status.Error(codes.InvalidArgument, "job_id and payload are required")
	}
	log.Printf("SubmitJob: job_id=%s", req.JobId)

	// Write to Redis hash for tracking
	s.Rdb.HSet(ctx, "job:"+req.JobId, map[string]interface{}{
		"status":     "pending",
		"retry":      0,
		"updated_at": ctx.Value("timestamp"), // optional
	})

	// Also enqueue in stream
	err := s.Rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "jobs",
		Values: map[string]interface{}{
			"job_id":  req.JobId,
			"payload": req.Payload,
			"retry":   0,
		},
	}).Err()

	if err != nil {
		metrics.JobsFailedByType.WithLabelValues("SubmitJob").Inc()
		log.Printf("SubmitJob: redis error: %v", err)
		return nil, err
	}

	return &pb.JobResponse{Message: "Job submitted successfully"}, nil
}
