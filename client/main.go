package main

import (
	"log"

	pb "github.com/johnsunam/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreenServiceClient(conn)
	names := &pb.NameList{
		Names: []string{"John", "Doe", "Bob"},
	}
	// RPC request response
	callSayHello(client)
	// RPC server stream
	callSayHelloServerStream(client, names)
	// RPC client stream
	callSayHelloClientStream(client, names)
	// RPC bidirectional stream
	callSayHelloBidrectionalStreaming(client)
}
