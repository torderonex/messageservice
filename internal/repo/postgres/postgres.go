package postgres

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/torderonex/messageservice/internal/config"
	"log"
	"time"
)

func New(cfg config.PostgresDatabase) *sqlx.DB {
	flag.Parse()
	var db *sqlx.DB

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, "disable"))

	if err != nil || db.Ping() != nil {
		err = doWithRetries(cfg, &db)
		if err != nil {
			log.Fatalf("Failed to connect to the database after %d retries", cfg.MaxRetries)
		}
	}

	return db
}

func doWithRetries(cfg config.PostgresDatabase, db **sqlx.DB) error {
	var err error
	for i := 0; i < cfg.MaxRetries; i++ {
		*db, err = sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, "disable"))
		if err == nil {
			return nil
		}

		log.Printf("Failed to connect to the database: %v", err)
		if i < cfg.MaxRetries-1 {
			log.Printf("Retrying in %v seconds...", cfg.RetryInterval.Seconds())
			time.Sleep(cfg.RetryInterval)
		}
	}
	return err
}
