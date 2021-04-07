package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
)

func main() {
	db := connect()
	defer db.Close()

	hash, err := shorten(db, "http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	orig, err := get(db, hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Get by hash: ", orig)
}

func shorten(db *pg.DB, url string) (string, error) {
	hash, err := getNewHash(db)
	if err != nil {
		return "", err
	}

	url1 := &URL{
		Original: "http://google.com",
		Hash:     hash,
	}
	_, err = db.Model(url1).Insert()
	if err != nil {
		return "", err
	}

	return url1.Hash, nil
}

func getNewHash(db *pg.DB) (string, error) {
	var url URL
	err := db.Model(&url).Where("original is null").Limit(1).Select()
	if err != nil {
		return generate(db)
	}
	return url.Hash, nil
}

var i int = 100000

func generate(db *pg.DB) (string, error) {
	hash := ""
	for count := 0; count < 5; count++ {
		i++
		hash = fmt.Sprintf("%d", i)
		url1 := &URL{
			Hash: hash,
		}
		_, err := db.Model(url1).Insert()
		if err != nil {
			return "", err
		}
	}

	return hash, nil
}

func get(db *pg.DB, hash string) (string, error) {
	var url URL

	var urls []URL
	err := db.Model(&urls).Select()
	if err != nil {
		return "", err
	}

	err = db.Model(&url).Where("hash = ? and original is not null", hash).Limit(1).Select()
	if err != nil {
		return "", err
	}
	return url.Original, nil
}
