package main

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	pb "github.com/jeisberg/go_protobuf_users/users"
)

const (
	address     = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	r, err := c.GetUser(ctx, &pb.UserMessage{Id: 5})
	
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	
	log.Printf("User: %s", r)
}
