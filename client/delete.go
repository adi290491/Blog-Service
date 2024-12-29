package main

import (
	"context"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was called---")

	data := &pb.BlogId{
		BlogId: id,
	}

	_, err := c.DeleteBlog(context.Background(), data)

	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog with id %s was deleted\n", id)
}
