package main

import (
	"io"
	"log"

	pb "github.com/johnsunam/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreenService_SayHelloClientStreamingServer) error {
	names := []string{}
	log.Printf("Received client stream request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: names})
		}
		if err != nil {
			return err
		}
		log.Println("streamed name: ", req.Name)
		names = append(names, "Hello ", req.Name)
	}
}
