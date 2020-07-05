package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"google.golang.org/grpc"
)

type readBlogCmdParams struct {
	BlogID *string
}

var readBlogParams *readBlogCmdParams

func init() {
	readBlogParams = &readBlogCmdParams{}

	readBlogParams.BlogID = readCmd.Flags().StringP("blogid", "r", "", "Id of the blog post you want to read")
	readCmd.MarkFlagRequired("blogid")

	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reading a blog post from the database",
	Long:  `Reading a blog post from the database`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer s.Close()

		c := pb.NewBlogServiceClient(s)

		blog, err := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{
			BlogId: *readBlogParams.BlogID,
		})

		if err != nil {
			log.Fatalf("something went wrong: %v", err)
		}

		fmt.Println(blog.Blog)
	},
}
