package db

import (
	"GamificationEducation/internal/config"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/pressly/goose"
	"log"
)

func NewDB(cfg *config.DB) (*sql.DB, error) {
	connCfg, err := pgx.ParseURI(cfg.URI)
	if err != nil {
		return nil, fmt.Errorf("parse DB URI: %w", err)
	}

	db := stdlib.OpenDB(connCfg)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("check DB connection: %w", err)
	}

	// make migrations

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "internal/db/migrations"); err != nil {
		panic(err)
	}
	log.Print("after migrations")

	return db, nil
}
