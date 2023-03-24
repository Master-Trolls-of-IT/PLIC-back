package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_unix.go: %s environment variable not set.\n", k)
		}
		return v
	}
	var (
		dbUser         = mustGetenv("my-db-user")     // e.g. 'my-db-user'
		dbPwd          = mustGetenv("my-db-password") // e.g. 'my-db-password'
		unixSocketPath = mustGetenv("instance")       // e.g. '/cloudsql/project:region:instance'
		dbName         = mustGetenv("my-database")    // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		dbUser, dbPwd, dbName, unixSocketPath)

	// db is the pool of database connections.
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return &PostgresDB{db}, nil
}
