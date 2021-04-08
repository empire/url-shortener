package hash

import (
	"context"
	"fmt"
	"time"

	pb "github.com/empire/url-shortener/api/hashgen"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func Generate() (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", fmt.Errorf("grpc.Dial: %w", err)
	}
	defer conn.Close()
	c := pb.NewHashGeneratorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Generate(ctx, &pb.HashRequest{})
	if err != nil {
		return "", fmt.Errorf("c.Generate: %w", err)
	}
	return r.GetHash(), nil
}
