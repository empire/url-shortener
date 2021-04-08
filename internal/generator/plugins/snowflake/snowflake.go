package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/empire/url-shortener/internal/unique"
)

type SnowFlake struct {
	node *snowflake.Node
}

func (s *SnowFlake) Generate() int64 {
	return s.node.Generate().Int64()
}

func MustNew() *SnowFlake {
	node, err := New()
	if err != nil {
		panic(err)
	}

	return node
}

func New() (*SnowFlake, error) {
	nodeid, err := unique.MachineId()
	if err != nil {
		return nil, err
	}

	node, err := snowflake.NewNode(int64(nodeid))
	if err != nil {
		return nil, err
	}
	return &SnowFlake{node}, nil
}
