package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	DB *sql.DB
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

	// db is the pool of database connections.

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return &Database{db}, nil
}
