package main

import (
	"context"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was called---")

	blog := &pb.Blog{
		Author:  "Aditya",
		Title:   "First Blog",
		Content: "My First Blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog created: %s\n", res.GetBlogId())
	return res.GetBlogId()
}
