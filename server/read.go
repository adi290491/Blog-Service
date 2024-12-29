package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Readblog was invoked with %v\n", req)

	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid blog ID: %s", req.GetBlogId()),
		)
	}

	data := &BlogItem{}

	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found with ID: %s", req.GetBlogId()),
		)
	}

	return documentToBlog(data), nil
}
