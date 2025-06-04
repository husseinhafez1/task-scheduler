package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"task/internal/server"
	"task/internal/worker"
	pb "task/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func connectRedisWithRetry() *redis.Client {
	var rdb *redis.Client
	var err error
	ctx := context.Background()

	for attempts := 0; attempts < 5; attempts++ {
		rdb = redis.NewClient(&redis.Options{
			Addr: "redis:6379", // use "redis:6379" if in Docker
		})

		err = rdb.Ping(ctx).Err()
		if err == nil {
			log.Println("Connected to Redis")
			return rdb
		}

		log.Printf("Failed to connect to Redis (attempt %d): %v", attempts+1, err)
		time.Sleep(time.Duration(attempts+1) * time.Second)
	}

	log.Fatalf("Could not connect to Redis after 5 attempts: %v", err)
	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rdb := connectRedisWithRetry()
	go worker.StartWorker(rdb)

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, server.NewServer(rdb))
	reflection.Register(grpcServer)

	go func() {
		log.Println("gRPC server running on port 50053")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	grpcServer.GracefulStop()
	log.Println("Server stopped.")
}
