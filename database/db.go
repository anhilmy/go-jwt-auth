package database

import (
	"fmt"
	"go-jwt-auth/internal/config"
	"log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

type Database struct {
	db *dbx.DB
}

func OpenDB(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("postgres://%s:%s/%s?sslmode=disable&user=%s&password=%s", config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)

	db := Database{}
	conn, err := dbx.MustOpen("postgres", dsn)

	// simple logging
	conn.LogFunc = log.Printf

	if err != nil {
		return nil, err
	}

	db.db = conn

	return &db, nil
}

func (db Database) DB() *dbx.DB {
	return db.db
}
