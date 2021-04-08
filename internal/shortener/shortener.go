package shortener

import (
	"log"
	"time"

	"github.com/empire/url-shortener/api/models"
	"github.com/go-pg/pg/v10"
)

type Shortener struct {
	db *pg.DB
}

func New() (*Shortener, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	return &Shortener{
		db: db,
	}, nil
}

func (s *Shortener) Close() {
	s.db.Close()
}

func (s *Shortener) Shorten(hash, original string, age int32) (string, error) {
	now := time.Now()
	url := &models.URL{
		Original:  original,
		Hash:      hash,
		CreatedAt: now,
		ExpiredAt: now.Add(time.Duration(age) * 24 * time.Hour),
	}
	// TODO add retries (it's rare case but it can duplciate the hash)
	rs, err := s.db.Model(url).Insert()
	log.Println(rs)
	if err != nil {
		return "", err
	}

	log.Println(url.String())
	return url.Hash, nil
}

func (s *Shortener) GetUrl(hash string) (string, error) {
	now := time.Now()
	var url models.URL

	err := s.db.Model(&url).Where("hash = ? and original is not null and expired_at > ?", hash, now).Limit(1).Select()
	if err != nil {
		return "", err
	}
	return url.Original, nil
}
