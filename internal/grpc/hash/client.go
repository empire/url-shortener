package hash

import (
	"context"
	"log"
	"time"

	pb "github.com/empire/url-shortener/api/hashgen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func Get() (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHashGeneratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Generate(ctx, &pb.HashRequest{})
	if err != nil {
		return "", err
	}
	return r.GetHash(), nil
}
