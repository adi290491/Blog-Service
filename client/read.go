package main

import (
	"context"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was called---")

	req := &pb.BlogId{
		BlogId: id,
	}

	blog, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	log.Printf("Blog found: %v\n", blog)
	return blog
}
