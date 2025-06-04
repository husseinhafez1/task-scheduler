package main

import (
	"log"
	"net"

	"task/internal/server"
	pb "task/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(grpcServer)
	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
