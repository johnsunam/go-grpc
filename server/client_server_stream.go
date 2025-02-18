package main

import (
	"io"
	"log"

	pb "github.com/johnsunam/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreenService_SayHelloBidirectionalStreamingServer) error {
	log.Printf("Received bidirectional stream request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("streamed name: ", req.Name)
		if err := stream.Send(&pb.HelloResponse{Message: req.Name}); err != nil {
			return err
		}
	}
}
