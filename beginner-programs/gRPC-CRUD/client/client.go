package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/client/cmd"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"google.golang.org/grpc"
)

func main() {
	cmd.Execute()

	s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Client started")
	}

	defer s.Close()

	c := pb.NewBlogServiceClient(s)

	request := &pb.Blog{
		AuthorId: "Gabriel",
		Title:    "GRCP CRUD tutorial",
		Content:  "Something about GRPC CRUD functionality",
	}

	createdBlog, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{
		Blog: request,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Blog has been created: %v", createdBlog.Blog)
	fmt.Println("")

	blog, err := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{
		BlogId: createdBlog.Blog.Id,
	})

	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	fmt.Println(blog.Blog)

	updateRequest := &pb.UpdateBlogRequest{
		BlogId: createdBlog.Blog.Id,
		Blog: &pb.Blog{
			AuthorId: "Gabriel123",
			Title:    "GRCP CRUD tutorial 123",
			Content:  "Something about GRPC CRUD functionality 123",
		},
	}

	updatedBlog, err := c.UpdateBlog(context.Background(), updateRequest)

	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	fmt.Println(updatedBlog.Blog)

	deleteResponse, err := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{
		BlogId: createdBlog.Blog.Id,
	})

	if err != nil {
		log.Fatalf("something went wrong: %v", err)
	}

	fmt.Println(deleteResponse.Status)
}
