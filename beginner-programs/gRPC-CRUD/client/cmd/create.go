package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tannergabriel/learning-go/beginner-programs/gRPC-CRUD/pb"
	"google.golang.org/grpc"
)

type createBlogCmdParams struct {
	ID       *string
	Title    *string
	AuthorID *string
	Content  *string
}

var createBlogParams *createBlogCmdParams

func init() {
	createBlogParams = &createBlogCmdParams{}

	createBlogParams.Title = createCmd.Flags().StringP("title", "t", "", "Title of the blog post")
	createCmd.MarkFlagRequired("title")

	createBlogParams.AuthorID = createCmd.Flags().StringP("authorid", "i", "", "Id of the blog post author")
	createCmd.MarkFlagRequired("authorid")

	createBlogParams.Content = createCmd.Flags().StringP("content", "c", "", "Blog post content")
	createCmd.MarkFlagRequired("content")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creating a new Blog post in the database",
	Long:  `Creating a new Blog post in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer s.Close()

		c := pb.NewBlogServiceClient(s)

		request := &pb.Blog{
			AuthorId: *createBlogParams.AuthorID,
			Title:    *createBlogParams.Title,
			Content:  *createBlogParams.Content,
		}

		createdBlog, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{
			Blog: request,
		})

		if err != nil {
			panic(err)
		}

		fmt.Printf("Blog has been created: %v", createdBlog.Blog)
	},
}
