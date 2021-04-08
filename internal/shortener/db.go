package shortener

import (
	"github.com/empire/url-shortener/api/models"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func connect() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.URL)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
