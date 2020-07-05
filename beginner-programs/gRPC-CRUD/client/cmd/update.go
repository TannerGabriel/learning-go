package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"google.golang.org/grpc"
)

type updateBlogCmdParams struct {
	BlogID   *string
	Title    *string
	AuthorID *string
	Content  *string
}

var updateBlogParams *updateBlogCmdParams

func init() {
	updateBlogParams = &updateBlogCmdParams{}

	updateBlogParams.BlogID = updateCmd.Flags().StringP("id", "", "", "Blog post id")
	updateCmd.MarkFlagRequired("id")

	updateBlogParams.Title = updateCmd.Flags().StringP("title", "t", "", "Title of the blog post")
	updateCmd.MarkFlagRequired("title")

	updateBlogParams.AuthorID = updateCmd.Flags().StringP("authorid", "i", "", "Id of the blog post author")
	updateCmd.MarkFlagRequired("authorid")

	updateBlogParams.Content = updateCmd.Flags().StringP("content", "c", "", "Blog post content")
	updateCmd.MarkFlagRequired("content")

	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an already existing Blog post in the database",
	Long:  `Update an already existing Blog post in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer s.Close()

		c := pb.NewBlogServiceClient(s)

		updateRequest := &pb.UpdateBlogRequest{
			BlogId: *updateBlogParams.BlogID,
			Blog: &pb.Blog{
				AuthorId: *updateBlogParams.AuthorID,
				Title:    *updateBlogParams.Title,
				Content:  *updateBlogParams.Content,
			},
		}

		updatedBlog, err := c.UpdateBlog(context.Background(), updateRequest)

		if err != nil {
			log.Fatalf("something went wrong: %v", err)
		}

		fmt.Println(updatedBlog.Blog)
	},
}
