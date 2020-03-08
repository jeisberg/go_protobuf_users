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

// SayHello implements helloworld.GreeterServer
func (s *server) GetUser(ctx context.Context, in *pb.UserMessage) (*pb.UserMessage, error) {
	
	log.Printf("Received: %v", in.GetId())
	
	return &pb.UserMessage {
		Id: in.GetId(),
		Name: "Name: Captain Funk",
		Email: "captainfunk@legit.com",
	}, 
	nil
}

func main() {
	
	lis, err := net.Listen("tcp", port)
	
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	
	pb.RegisterUserServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
