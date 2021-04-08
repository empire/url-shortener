package generator

import (
	"github.com/empire/url-shortener/internal/encoding"
	"github.com/empire/url-shortener/internal/generator/plugins/snowflake"
)

var counter int64
var generator interface{ Generate() int64 }

func init() {
	generator = snowflake.MustNew()
}

func New() string {
	return encoding.Encode(generator.Generate())
}
