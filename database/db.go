package database

import (
	"fmt"
	"log"
	"os"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	DbName     string `yaml:"db_name"`
	DbPort     string `yaml:"db_port"`
	DbHost     string `yaml:"db_host"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
}

func ReadConf(filename string) (*DbConfig, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	res := DbConfig{}
	err = yaml.Unmarshal(buf, &res)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return &res, err
}

type Database struct {
	db *dbx.DB
}

func OpenDB(config *DbConfig) (*Database, error) {
	dsn := fmt.Sprintf("postgres://%s:%s/%s?sslmode=disable&user=%s&password=%s", config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	log.Println(dsn)

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
