package generator

import (
	"github.com/empire/url-shortener/internal/encoding"
	"github.com/empire/url-shortener/internal/generator/plugins/snowflake"
)

var counter int64
var generator interface{ Generate() int64 }

// First 20bits are redundant(in most cases), we can use shorter code
var Mask = int64(0x00_00_0F_FF_FF_FF_FF_FF)

func init() {
	generator = snowflake.MustNew()
}

func New() string {
	return encoding.Encode(generator.Generate() & Mask)
}
