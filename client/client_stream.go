package main

import (
	"context"
	"log"
	"time"

	pb "github.com/johnsunam/go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreenServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to call SayHelloClientStreaming: %v", err)
	}
	log.Println("******** Started client streaming ********")
	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("failed to send request: %v", err)
		}
		time.Sleep(3 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}
	log.Printf("Response from server: %s", resp.Message)
}
