package config

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type Scripts map[string]string

type DB struct {
	Client  *sqlx.DB
	Scripts *Scripts
}

func BootDatabase(dbConfig DBConfig) (*DB, error) {
	db, err := openSqlxViaPgxConnPool(dbConfig)
	scripts, errScripts := LoadSQLScripts()

	if err != nil {
		return nil, err
	}
	if errScripts != nil {
		return nil, err
	}

	return &DB{
		Client:  db,
		Scripts: scripts,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

// SQLX connection
// func get(connStr string) (*sqlx.DB, error) {
// 	db, err := sqlx.Connect("postgres", connStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetMaxOpenConns(180)
// 	db.SetMaxIdleConns(20)
// 	db.SetConnMaxLifetime(2 * time.Minute)

// 	if err := db.Ping(); err != nil {
// 		return nil, err
// 	}

// 	return db, nil
// }

// PGX Connection
func openSqlxViaPgxConnPool(dbConfig DBConfig) (*sqlx.DB, error) {
	port, _ := strconv.ParseUint(dbConfig.Port, 10, 16)
	connConfig := pgx.ConnConfig{
		Host:     dbConfig.Host,
		Port:     uint16(port),
		User:     dbConfig.User,
		Password: dbConfig.Pwd,
		Database: dbConfig.Name,
	}
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("call to pgx.NewConnPool failed: %w", err)
	}

	// Apply any migrations...

	// Then set up sqlx and return the created DB reference
	nativeDB := stdlib.OpenDBFromPool(connPool)
	return sqlx.NewDb(nativeDB, "pgx"), nil
}

func (s *Scripts) Get(script string) string {
	return (*s)[script]
}

const dir = "scripts"

//go:embed scripts/*
var scriptsFS embed.FS

// LOAD ALL FILES
func LoadSQLScripts() (*Scripts, error) {
	scripts := make(Scripts)

	files, _ := fs.ReadDir(scriptsFS, dir)
	for _, script := range files {
		name := script.Name()
		cleanName := strings.Split(name, ".")[0]

		if len(scripts[cleanName]) > 0 {
			log.Error().Msgf("Script name collision")
			return nil, errors.New("script name collision")
		}

		file, _ := scriptsFS.ReadFile(dir + "/" + name)
		scripts[cleanName] = string(file)
	}

	return &scripts, nil
}
