package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/jeisberg/go_protobuf_users/users"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// implementation of our goofy GetUser function
func (s *server) GetUser(ctx context.Context, in *pb.UserMessage) (*pb.UserMessage, error) {
	// log the message here
	log.Printf("Received: %v", in.GetId())
	// return a user
	return &pb.UserMessage {
		Id: in.GetId(),
		Name: "Captain Funk",
		Email: "captainfunk@legit.com",
	}, 
	nil // no error for now
}

func main() {
	// listen on the port
	lis, err := net.Listen("tcp", port)
	// sanity check
	if err != nil {
		// fatal didn't bind
		log.Fatalf("failed to listen: %v", err)
	}
	// new server
	s := grpc.NewServer()
	// register the user service
	pb.RegisterUserServiceServer(s, &server{})
	// sanity check that we're listening
	if err := s.Serve(lis); err != nil {
		// fatal didn't serve
		log.Fatalf("failed to serve: %v", err)
	}
}
