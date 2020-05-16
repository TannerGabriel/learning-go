package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-Example/chat"
)

func main() {
	// Create a client gRPC connection
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// Create a chat service client using the connection
	c := chat.NewChatServiceClient(conn)

	// Call the SayHello() function of the service
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
