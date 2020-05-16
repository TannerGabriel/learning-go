package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-Example/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server listening on port 9000")

	// Create chat service
	s := chat.Server{}

	// Create new gRPC server
	grpcServer := grpc.NewServer()

	// Register the chat service to the gRPC server
	chat.RegisterChatServiceServer(grpcServer, &s)

	// Serve the gRPC server to a client using net
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
