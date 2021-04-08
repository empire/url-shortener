package models

import (
	"fmt"
	"time"
)

type URL struct {
	Id        int64
	Original  string `pg:",unique"`
	Hash      string `pg:",unique"`
	CreatedAt time.Time
	ExpiredAt time.Time
}

func (u URL) String() string {
	return fmt.Sprintf("URL<%d %s %s %s %s>", u.Id, u.Original, u.Hash, u.CreatedAt, u.ExpiredAt)
}
