package models

import (
	"fmt"
	"time"

	_ "github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
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
