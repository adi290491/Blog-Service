package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:9000"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on %v\n", err)
	}
}
