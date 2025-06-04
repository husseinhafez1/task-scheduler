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

func (s *TaskServer) SubmitJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	err := s.Rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "jobs",
		Values: map[string]interface{}{
			"job_id":  req.JobId,
			"payload": req.Payload,
			"retry":   0,
		},
	}).Err()

	if err != nil {
		return nil, err
	}

	return &pb.JobResponse{Message: "Job submitted successfully"}, nil
}
