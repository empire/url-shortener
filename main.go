package main

import (
	"context"
	"log"

	"github.com/empire/url-shortener/internal/grpc/shortener"
	_ "github.com/empire/url-shortener/internal/grpc/shortener"
)

// TODO get custom hash
// Make production ready

func main() {
	ctx := context.Background()
	hash, err := shortener.Shorten(ctx, "http://google.com", 5)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hash)
	u, e := shortener.GetUrl(ctx, hash)
	log.Println(u, e)
}
