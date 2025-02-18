package main

import (
	"context"
	"log"

	pb "github.com/johnsunam/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Printf("Received unary request")
	return &pb.HelloResponse{Message: "Hello "}, nil
}
