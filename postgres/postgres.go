package postgres

import (
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Open(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", url)

	if err != nil {
		log.Fatal("Error connection database", err)
		return nil, err
	}

	return db, nil
}
