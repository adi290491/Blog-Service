package main

import (
	"log"

	pb "github.com/adi290491/go-grpc-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:9000"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect on %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// createBlog(c)

	//valid id
	// readBlog(c, "67146478454f7b1f3fde7508")

	//invalid id
	// readBlog(c, "67146478454f7b1f3fde7509")

	// updateBlog(c, "67146478454f7b1f3fde7508")

	// listBlogs(c)
	deleteBlog(c, "67146478454f7b1f3fde7508")

	//invalid delete
	deleteBlog(c, "1")

}
