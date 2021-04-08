package shortener

import (
	"context"
	"log"
	"time"

	pb "github.com/empire/url-shortener/api/shorten"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func Shorten(ctx context.Context, url string, age int32) (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShortenerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.Shorten(ctx, &pb.ShortenRequest{Url: url, Age: age})
	if err != nil {
		return "", err
	}
	return r.GetCode(), nil
}

func GetUrl(ctx context.Context, code string) (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewShortenerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.GetUrl(ctx, &pb.GetUrlRequest{Code: code})
	if err != nil {
		return "", err
	}
	return r.GetUrl(), nil
}
