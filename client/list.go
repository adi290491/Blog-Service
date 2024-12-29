package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {
	log.Println("---ListBlogs was called---")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling listBlogs: %v\n", err)
		return
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		log.Println(res)
	}
}
