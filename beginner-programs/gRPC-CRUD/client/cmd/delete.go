package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"google.golang.org/grpc"
)

type deleteBlogCmdParams struct {
	BlogID *string
}

var deleteBlogParams *deleteBlogCmdParams

func init() {
	deleteBlogParams = &deleteBlogCmdParams{}

	deleteBlogParams.BlogID = deleteCmd.Flags().StringP("blogid", "r", "", "Id of the blog post you want to read")
	deleteCmd.MarkFlagRequired("blogid")

	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a blog post from the database",
	Long:  `Delete a blog post from the database`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer s.Close()

		c := pb.NewBlogServiceClient(s)

		deleteResponse, err := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{
			BlogId: *deleteBlogParams.BlogID,
		})

		if err != nil {
			log.Fatalf("something went wrong: %v", err)
		}

		fmt.Println(deleteResponse.Status)
	},
}
