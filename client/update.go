package main

import (
	"context"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {

	log.Println("---update was called---")
	// Update blog with provided ID
	newBlog := &pb.Blog{
		Id:      id,
		Author:  "Jerry",
		Title:   "Memories",
		Content: "Memories bring back you",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error while updating: %v\n", err)
	}

}
