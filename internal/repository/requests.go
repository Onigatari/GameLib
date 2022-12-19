package repository

import "github.com/jmoiron/sqlx"

type RequestPostgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *RequestPostgres {
	return &RequestPostgres{db: db}
}
