package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		log.Fatalf("error connect to DB\n")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error ping to DB %s %s:%s \n", err, cfg.Host, cfg.Port)
		return nil, err
	}

	return db, nil
}
