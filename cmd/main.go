package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"task/internal/server"
	"task/internal/worker"
	pb "task/proto"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func connectRedisWithRetry(addr string) *redis.Client {
	var rdb *redis.Client
	var err error
	ctx := context.Background()

	for attempts := 0; attempts < 5; attempts++ {
		rdb = redis.NewClient(&redis.Options{
			Addr: addr,
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
	// Configurable ports and addresses
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "redis:6379"
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50053"
	}
	metricsPort := os.Getenv("METRICS_PORT")
	if metricsPort == "" {
		metricsPort = "2113"
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rdb := connectRedisWithRetry(redisAddr)
	go worker.StartWorker(rdb)

	// Start Prometheus metrics server with /metrics and /healthz endpoints
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	metricsSrv := &http.Server{Addr: ":" + metricsPort, Handler: mux}
	go func() {
		log.Printf("Prometheus metrics server running on :%s", metricsPort)
		if err := metricsSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Metrics server failed: %v", err)
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, server.NewServer(rdb))
	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRPC server running on port %s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	grpcServer.GracefulStop()
	_ = metricsSrv.Shutdown(context.Background())
	log.Println("Server stopped.")
}
