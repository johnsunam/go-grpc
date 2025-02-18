package main

import (
	"context"
	"io"
	"log"

	pb "github.com/johnsunam/go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreenServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("failed to call SayHelloServerStreaming: %v", err)
	}
	log.Println("******** Started server streaming ********")
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Printf("Response from server: %s", resp.Message)
	}
}
