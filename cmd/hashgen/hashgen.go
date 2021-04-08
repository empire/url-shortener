package main

// Based on https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
import (
	"context"
	"log"
	"net"

	pb "github.com/empire/url-shortener/api/hashgen"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedHashGeneratorServer
}

func (s *server) Generate(ctx context.Context, in *pb.HashRequest) (*pb.HashReply, error) {
	return &pb.HashReply{Hash: "Hello"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listen to: %s\n", port)
	s := grpc.NewServer()
	pb.RegisterHashGeneratorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
