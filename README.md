# vtr-mailer-service

## How to Run
### Development
`docker-compose up`

### Production
`docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d`

### TODO
- [ ] Integrate with PostgreSQL
  - [x] Configure PGX as database driver
  - [x] Configure SQLX
  - [ ] Integrate a migration solution
- [ ] Integrate with another Email solution
- [ ] Webhooks from Email providers
- [ ] Webhooks to consumers
- [ ] API Auth
  - [ ] User
  - [ ] Authentication
  - [ ] Authorization
- [ ] Work with queue
  - [ ] Queue consumer

## External Libs
Framework Web: Gin Gonic - https://github.com/gin-gonic/gin
CORS: Gin Contrib / Cors - https://github.com/gin-contrib/cors
DotEnv: GoDotEnv - https://github.com/joho/godotenv
Logger: Zerolog - https://github.com/rs/zerolog
Migration: Goose - https://github.com/pressly/goose
PostgreSQL Driver: pgx - https://github.com/jackc/pgx
Database/SQL: sqlx - https://github.com/jmoiron/sqlx

### PGX/SQLX Example
Guide: https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4


```go
package main

import (
	"log"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func main() {
	_, err := OpenSqlxViaPgxConnPool()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Worked!")
	}
}

// OpenSqlxViaPgxConnPool does what its name implies
func OpenSqlxViaPgxConnPool() (*sqlx.DB, error) {
	// First set up the pgx connection pool
	connConfig := pgx.ConnConfig{
		Host:     "localhost",
		Database: "my_local_db",
	}
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Call to pgx.NewConnPool failed")
	}

	// Apply any migrations...

	// Then set up sqlx and return the created DB reference
	nativeDB, err := stdlib.OpenFromConnPool(connPool)
	if err != nil {
		connPool.Close()
		return nil, errors.Wrap(err, "Call to stdlib.OpenFromConnPool failed")
	}
	return sqlx.NewDb(nativeDB, "pgx"), nil
}
```