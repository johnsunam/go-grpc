package main

import (
	"log"
	"time"

	pb "github.com/johnsunam/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(names *pb.NameList, stream pb.GreenService_SayHelloServerStreamingServer) error {
	log.Printf("Received server stream request: %+v", names.Names)
	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloResponse{Message: "Hello " + name}); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
