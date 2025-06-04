package main

import (
	"context"
	"log"
	pb "task/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	res, err := client.SubmitJob(context.Background(), &pb.JobRequest{
		JobId:   "job456",
		Payload: "send email",
	})

	if err != nil {
		log.Fatalf("could not submit job: %v", err)
	}

	log.Println("Server response:", res.Message)
}
