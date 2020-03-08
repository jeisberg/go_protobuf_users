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
	// dial up the service endpoint here
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	// sanity check
	if err != nil {
		// fatal didn't bind here
		log.Fatalf("did not connect: %v", err)
	}
	// close on defer
	defer conn.Close()
	// new client here
	api := pb.NewUserServiceClient(conn)
	// set the timeout 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel
	defer cancel()
	// get the user from the client
	user, err := api.GetUser(ctx, &pb.UserMessage{Id: 5})
	// sanity check
	if err != nil {
		// could not get the user here
		log.Fatalf("could not get user: %v", err)
	}
	// print the user
	log.Printf("User: %s", user)
}
