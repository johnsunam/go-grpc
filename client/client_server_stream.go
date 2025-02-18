package main

import (
	"context"
	"log"
	"time"

	pb "github.com/johnsunam/go-grpc/proto"
)

func callSayHelloBidrectionalStreaming(client pb.GreenServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.SayHelloBidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalf("failed to call SayHelloBidrectionalStreaming: %v", err)
	}
	names := []string{"Lucy", "Mote", "langti", "puntu"}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("failed to send name: %v", err)
		}
		resp, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Printf("Bidirectional response from server: %s", resp.Message)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("failed to close stream: %v", err)
	}
}
