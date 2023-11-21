package repository

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase(URI ...string) (*Database, error) {
	var dbURI string
	if len(URI) > 0 {
		dbURI = URI[0]
	} else {
		mustGetenv := func(variable string) string {
			value := os.Getenv(variable)
			if value == "" {
				log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.\n", variable)
			}
			return value
		}
		var (
			dbUser         = mustGetenv("my-db-user")     // e.g. 'my-db-user'
			dbPwd          = mustGetenv("my-db-password") // e.g. 'my-db-password'
			unixSocketPath = mustGetenv("instance")       // e.g. '/cloudsql/project:region:instance'
			dbName         = mustGetenv("my-database")    // e.g. 'my-database'
		)

		dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s",
			dbUser, dbPwd, dbName, unixSocketPath)
	}

	db, err := sqlx.Connect("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect: %v", err)
	}

	return &Database{DB: db}, nil
}

func (d *Database) Get(dest interface{}, query string, args ...interface{}) error {
	return d.DB.Get(dest, query, args...)
}
