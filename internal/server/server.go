package server

import (
	"context"

	pb "task/proto"

	"github.com/go-redis/redis/v8"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	Rdb *redis.Client
}

func NewServer(rdb *redis.Client) *TaskServer {
	return &TaskServer{Rdb: rdb}
}

func (s *TaskServer) GetJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	status, err := s.Rdb.HGet(ctx, "job:"+req.JobId, "status").Result()
	if err == redis.Nil {
		return &pb.JobStatusResponse{Status: "not found"}, nil
	} else if err != nil {
		return nil, err
	}
	return &pb.JobStatusResponse{Status: status}, nil
}

func (s *TaskServer) SubmitJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	jobID := req.JobId

	// Write to Redis hash for tracking
	s.Rdb.HSet(ctx, "job:"+jobID, map[string]interface{}{
		"status":     "pending",
		"retry":      0,
		"updated_at": ctx.Value("timestamp"), // optional
	})

	// Also enqueue in stream
	err := s.Rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "jobs",
		Values: map[string]interface{}{
			"job_id":  jobID,
			"payload": req.Payload,
			"retry":   0,
		},
	}).Err()

	if err != nil {
		return nil, err
	}

	return &pb.JobResponse{Message: "Job submitted successfully"}, nil
}
