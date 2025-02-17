package main

import (
	"context"
	"log"
	"time"

	pb "github.com/johnsunam/go-grpc/proto"
)

func callSayHello(client pb.GreenServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("Response from server: %s", resp.Message)
}
