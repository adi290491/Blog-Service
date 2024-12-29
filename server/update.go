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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, req *pb.Blog) (*emptypb.Empty, error) {

	log.Printf("UpdateBlog was called with %v\n", req)

	oid, err := primitive.ObjectIDFromHex(req.GetId())

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid blog ID: %s", req.GetId()),
		)
	}

	data := &BlogItem{
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not update: %v\n", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found with ID: %s", req.GetId()),
		)
	}

	log.Println("Blog was updated")
	return &emptypb.Empty{}, nil
}
