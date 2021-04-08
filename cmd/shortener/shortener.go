package main

import (
	"context"
	"log"
	"net"

	pb "github.com/empire/url-shortener/api/shorten"
	hashgrpc "github.com/empire/url-shortener/internal/grpc/hash"
	"github.com/empire/url-shortener/internal/shortener"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	shortener *shortener.Shortener
	pb.UnimplementedShortenerServer
}

func (s *server) Shorten(ctx context.Context, in *pb.ShortenRequest) (*pb.ShortenReply, error) {
	hash, err := getOrGenHash(in)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	code, err := s.shortener.Shorten(hash, in.GetUrl(), in.GetAge())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.ShortenReply{Code: code}, nil
}

func getOrGenHash(in *pb.ShortenRequest) (string, error) {
	hash := in.GetHash()
	if hash != "" {
		return hash, nil
	}
	return hashgrpc.Generate()
}

func (s *server) GetUrl(ctx context.Context, in *pb.GetUrlRequest) (*pb.GetUrlReply, error) {
	url, err := s.shortener.GetUrl(in.GetCode())
	if err != nil {
		return nil, err
	}

	return &pb.GetUrlReply{Url: url}, nil
}

func main() {
	shortener, err := shortener.New()
	if err != nil {
		log.Fatal(err)
	}
	defer shortener.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	log.Printf("Listen to: %s\n", port)
	s := grpc.NewServer()
	pb.RegisterShortenerServer(s, &server{shortener: shortener})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
