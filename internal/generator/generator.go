package generator

import (
	"log"

	"github.com/empire/url-shortener/internal/encoding"
	"github.com/empire/url-shortener/internal/generator/plugins/snowflake"
)

var counter int64

var generator interface{ Generate() int64 }

// First 20bits are redundant(in most cases), we can use shorter code
var Mask = int64(0x00_00_0F_FF_FF_FF_FF_FF)

// For concurrent use safety
var codes chan string

func init() {
	generator = snowflake.MustNew()
	codes = make(chan string)
	go func() {
		// TODO the code is error prune
		for {
			codes <- encoding.Encode(generator.Generate() & Mask)
		}
	}()
}

func GenNewCode() string {
	code := <-codes
	log.Printf("New code is generated: %s\n", code)
	return code
}
