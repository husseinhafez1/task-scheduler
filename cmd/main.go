package main

import (
	"log"
	"net"

	"github.com/husseinhafez1/task-scheduler/internal/server"
	pb "github.com/husseinhafez1/task-scheduler/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, server.NewServer(rdb))

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
